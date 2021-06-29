package main

import (
	"fmt"
)

// f is right channel add one send left channel
func f(left, right chan int) {
	left <- 1 + <-right
}

func f2(right chan int) {
	right <- 1 + <-right
}

const n = 10

// todo
// code source:
// https://talks.golang.org/2012/concurrency.slide#39
func main() {
	leftmost := make(chan int)
	left := leftmost
	right := leftmost

	for i := 0; i < n; i++ {
		right = make(chan int)
		// the first goroutine holds(leftmost, new chan)
		// the second goroutine holds(last right chan, new chan)
		go f(left, right)

		// 记录上一次right goroutine的值
		left = right
	}

	go func(right chan int) {
		right <- 1
	}(right)

	fmt.Println(<-leftmost)
}
