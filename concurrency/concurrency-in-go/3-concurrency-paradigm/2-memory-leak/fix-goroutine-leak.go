package main

import (
	"fmt"
	"time"
)

func main() {

	// done用来控制协程的退出
	doWork := func(done <-chan interface{}, str <-chan string) <-chan interface{} {
		terminated := make(chan interface{})
		go func() {
			defer fmt.Println("doWork exited.")
			defer close(terminated)

			for {
				select {
				case s := <-str:
					fmt.Printf("do work %s\n", s)
				case <-done:	//2. 控制协程的退出，避免出现协程泄露
					return
				}
			}
		}()
		return terminated
	}

	done := make(chan interface{})
	terminated := doWork(done, nil)

	go func() {
		// cancel the operation after 1s
		time.Sleep(1*time.Second)
		fmt.Println("canceling doWork goroutine...")
		close(done)
	}()

	<-terminated
	fmt.Println("done, main exited")
}
