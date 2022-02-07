package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)

	for i := 0; i < 10; i++ {
		go func(n int) {
			fmt.Println(n)
		}(i)
	}
	for {

	}
}
