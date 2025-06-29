package main

func main() {

}

type MapResp struct {
	value  int // 值类型固定为 interface{}
	ok     bool
	length int
}
type MapReq struct {
	op       string
	key, val int
	respChan chan MapResp
}

type SyncMap struct {
	reqChan  chan MapReq
	stopChan chan struct{}
}

func NewSyncMap() *SyncMap {
	s := &SyncMap{
		reqChan:  make(chan MapReq),
		stopChan: make(chan struct{}),
	}
	go func() {}()
	return s
}

func (m *SyncMap) run() {
	data := map[int]int{}
	for {
		select {
		case req := <-m.reqChan:
			switch req.op {
			case "get":
				val, ok := data[req.key]
				req.respChan <- MapResp{value: val, ok: ok}
			case "set":
				data[req.key] = req.val
				close(req.respChan)
			case "delete":
				delete(data, req.key)
				close(req.respChan)
			case "len":
				req.respChan <- MapResp{length: len(data)}
			}

		case <-m.stopChan:
			return
		}
	}
}

func (m *SyncMap) Set(key, val int) {
	respChan := make(chan MapResp, 1)
	m.reqChan <- MapReq{"set", key, val, respChan}
	<-respChan
}

func (m *SyncMap) Get(key int) (int, bool) {
	respChan := make(chan MapResp, 1)
	m.reqChan <- MapReq{op: "get", key: key, respChan: respChan}
	resp := <-respChan
	return resp.value, resp.ok
}

func (m *SyncMap) Delete(key int) {
	respChan := make(chan MapResp, 1)
	m.reqChan <- MapReq{op: "delete", key: key, respChan: respChan}
	<-respChan
}

func (m *SyncMap) Len() int {
	respChan := make(chan MapResp, 1)
	m.reqChan <- MapReq{op: "len", respChan: respChan}
	resp := <-respChan
	return resp.length
}

func (m *SyncMap) Stop() {
	m.stopChan <- struct{}{}
}
