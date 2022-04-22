package main

import (
	"net/http"
	"runtime"
	"sync"

	_ "net/http/pprof"
)

func init() {
	runtime.SetMutexProfileFraction(1)
}

func main() {
	m := sync.Mutex{}
	datas := make(map[int]struct{})

	for i := 0; i < 999; i++ {
		go func(i int) {
			m.Lock()
			defer m.Unlock()
			datas[i] = struct{}{}
		}(i)
	}

	_ = http.ListenAndServe(":6061", nil)
}
