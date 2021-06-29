package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boring(msg string) <-chan string {
	ch := make(chan string)
	go func() {
		for i := 0; ; i++ {
			ch <- fmt.Sprintf("%s --> %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return ch
}

// for-select pattern
// code source:
// https://talks.golang.org/2012/concurrency.slide#35
func main() {
	ch := boring("Joe")

	timeout := time.After(3 * time.Second)
	for {
		select {

		case s := <-ch:
			fmt.Println(s)

		case <-time.After(500 * time.Millisecond): // timeout using select
			fmt.Println("Timeout, you're too slow")
			return

		case <-timeout: // timeout for whole conversation using select
			fmt.Println("Timeout, you're talk too much")
			return
		}
	}
}
