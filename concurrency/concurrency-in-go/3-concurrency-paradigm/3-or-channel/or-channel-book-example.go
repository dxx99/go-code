package main

import (
	"fmt"
	"time"
)

func main() {
	var or func(channels ...<-chan interface{}) <-chan interface{}
	or = func(channels ...<-chan interface{}) <-chan interface{} { // 接收数量可变的通道，并返回单通道
		switch len(channels) {
		case 0:
			return nil // 递归函数，必须要有终止条件，如果是空切片，返回nil通道，这与不传通道的想法一致
		case 1:
			return channels[0] //如果切片只有一个元素，就返回
		}

		orDone := make(chan interface{})
		go func() { //建立goroutine 以并消息没有被阻塞
			defer close(orDone)
			switch len(channels) {
			case 2: // 由于是递归调用，每次递归调用至少有两个通道
				select {
				case <-channels[0]:
				case <-channels[1]:
				}
			default:
				select {
				case <-channels[0]:
				case <-channels[1]:
				case <-channels[2]:
				case <-or(append(channels[3:], orDone)...):
				}
			}
		}()
		return orDone
	}

	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or(sig(2*time.Hour), sig(5*time.Minute), sig(1*time.Second), sig(1*time.Minute))
	fmt.Printf("done after %v\n", time.Since(start))
}
