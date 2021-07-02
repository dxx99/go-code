package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	//runtime.GOMAXPROCS(1)
	runtime.GOMAXPROCS(2)
	a := 0
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				a += 1
			}
		}()
	}
	wg.Wait()
	fmt.Println(a)
}
