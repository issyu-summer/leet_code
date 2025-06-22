package main

func main() {

}

type waitEntry struct {
	val   int
	ch    chan struct{} // 多个等待者共享一个 channel
	isSet bool
}
type SimpleBlockingMap struct {
	store map[int]*waitEntry
}

func NewSimpleBlockingMap() *SimpleBlockingMap {
	return &SimpleBlockingMap{
		store: make(map[int]*waitEntry),
	}
}

func (bm *SimpleBlockingMap) Get(key int) int {
	elem, ok := bm.store[key]
	if ok {
		if elem.isSet {
			val := elem.val
			return val
		}
		ch := elem.ch
		<-ch // 等待 Set 的 close()
		return elem.val
	}

	// 第一个等待者创建 ch，其它 get 也会复用
	ch := make(chan struct{})
	bm.store[key] = &waitEntry{
		ch: ch,
	}

	<-ch
	val := bm.store[key].val
	return val
}

func (bm *SimpleBlockingMap) Set(key, val int) {
	elem, ok := bm.store[key]
	if !ok {
		bm.store[key] = &waitEntry{
			val:   val,
			isSet: true,
		}
		return
	}
	if elem.isSet {
		return
	}
	elem.val = val
	elem.isSet = true
	close(elem.ch) // ✨ 唤醒所有等待者
}
