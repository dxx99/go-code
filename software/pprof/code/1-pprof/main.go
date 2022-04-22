package main

import (
	"1-pprof/lib"
	"log"
	"net/http"
	"time"

	// 添加pprof
	_ "net/http/pprof"
)





func main() {
	go func() {
		for {
			log.Printf("len: %d", lib.Add("go-programing-tour-book"))
			time.Sleep(time.Millisecond * 10)
		}
	}()

	_ = http.ListenAndServe("0.0.0.0:6060", nil)
}


