package main

import (
	"context"
	"log"
	"sync"
	"time"
)

// code source:
// https://www.youtube.com/watch?v=LSzR0VEraWw
func sleepAndTalk(ctx context.Context, d time.Duration, msg string) {
	select {
	case <-time.After(d):
		log.Println(msg)
	case <-ctx.Done():
		log.Println(ctx.Err())
	}

}

var wg sync.WaitGroup

// use context control goroutine exit
func main() {
	log.Println("starting...")
	ctx, cancel := context.WithCancel(context.Background())

	// 一秒之后执行这个cancel方法
	time.AfterFunc(time.Second, cancel)

	wg.Add(1)
	go func() {
		defer wg.Done()
		sleepAndTalk(ctx, 5*time.Second, "hello")
	}()
	wg.Wait()
}
