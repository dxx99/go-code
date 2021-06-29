package main

import (
	"fmt"
	"time"
)

func main() {

	orChannel := func(channels ...<-chan interface{}) <-chan interface{} {
		done := make(chan interface{})
		for _, item := range channels {
			go func(ch <-chan interface{}) {
				select {
				case <-ch:
					close(done)
				case <-done: //防止协程泄露
					fmt.Printf("协程退出， %v\n", item)
					return
				}
			}(item)
		}
		return done
	}

	timeWork := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	startTime := time.Now()
	<-orChannel(timeWork(1*time.Hour), timeWork(3*time.Minute), timeWork(2*time.Second), timeWork(3*time.Hour))
	fmt.Printf("work time after %v\n", time.Since(startTime))

	time.Sleep(1 * time.Second)
}
