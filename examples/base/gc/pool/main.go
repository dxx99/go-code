package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"sync"
)

func main() {
	go func() {
		http.ListenAndServe(":6060", nil)
	}()

	http.HandleFunc("/example", func(writer http.ResponseWriter, request *http.Request) {
		//b := newBuf()
		b := bufPool.Get().([]byte)	// 优化

		// 模拟一些工作
		for idx := range b {
			b[idx] = 1
		}

		fmt.Fprintf(writer, "done %v", request.URL.Path[1:])

		// 放回到内存池中
		bufPool.Put(b)
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

var bufPool = sync.Pool{New: func() interface{} {
	return make([]byte, 10<<20)
}}

func newBuf() []byte {
	return make([]byte, 10<<20)
}
