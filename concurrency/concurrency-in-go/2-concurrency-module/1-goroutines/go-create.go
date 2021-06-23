package main

import (
	"fmt"
	"time"
)

func main() {
	num := 1
	for {
		go func() {
			time.Sleep(10 * time.Second)
		}()
		num++
		if num > 10000 {
			fmt.Println("goroutine: ", num)
		}
	}
}
