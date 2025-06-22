package main

import "context"

func main() {

}

type waiter struct {
	ch chan interface{}
}

type entry struct {
	val      interface{}
	isSet    bool
	waitList []*waiter
}

type BlockingMap struct {
	store map[interface{}]*entry
}

func NewBlockingMap() *BlockingMap {
	return &BlockingMap{
		store: make(map[interface{}]*entry),
	}
}

func (bm *BlockingMap) Get(ctx context.Context, key interface{}) (interface{}, error) {
	e, ok := bm.store[key]
	if !ok {
		e = &entry{}
		bm.store[key] = e
	}
	if e.isSet {
		val := e.val
		return val, nil
	}

	w := &waiter{ch: make(chan interface{}, 1)}
	e.waitList = append(e.waitList, w)

	select {
	case val := <-w.ch:
		return val, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func (bm *BlockingMap) Set(key, val interface{}) {
	e, ok := bm.store[key]
	if !ok {
		e = &entry{
			val:   val,
			isSet: true,
		}
		bm.store[key] = e
		return
	}
	if e.isSet {
		return
	}
	e.val = val
	e.isSet = true
	// 通知所有 waiters
	for _, w := range e.waitList {
		w.ch <- val
	}
	e.waitList = nil
}
