package main

import (
	"fmt"
	"sync"
	"time"
)

// 适用场景: 不关心并发结果， 或者有其他方式收集结果
// 把WaitGroup视作一个安全的并发计数器：调用Add增加计数，调用Done减少计数。调用Wait会阻塞并等待至计数器归零。
func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("1st goroutine sleeping...")
		time.Sleep( 1*time.Second )
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("2nd goroutine sleeping...")
		time.Sleep(2*time.Second)
	}()

	wg.Wait()
	fmt.Println("all goroutine finish.")

}
