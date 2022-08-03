package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/trace"
	"time"
)

func main() {
	f, _ := os.Create("./base/gc/trace.out")
	defer f.Close()
	trace.Start(f)
	defer trace.Stop()

	keepAlloc()
	keepAllocV2()

	go func() {
		for true {
			t := time.Now()
			runtime.GC()
			fmt.Println(time.Since(t))
			time.Sleep(1*time.Second)
		}
	}()

	time.Sleep(10*time.Second)

}

func allocate() {
	for i := 0; i < 100000; i++ {
		_ = make([]byte, 1<<20)
	}
}


var cache = map[interface{}]interface{}{}

func keepAlloc() {
	for i := 0; i < 10000; i++ {
		m := make([]byte, 1<<10)
		cache[i] = m
	}
}

func keepAllocV2() {
	for i := 0; i < 100000; i++ {
		go func() {
			select {
			}
		}()
	}
}