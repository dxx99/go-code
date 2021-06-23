package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	var shareLock sync.Mutex
	const rTime = 1 * time.Second

	// 贪婪工作模式，整个工作循环中共享锁
	greedyWorker := func() {
		defer wg.Done()
		count := 0
		// 1秒钟执行的任务
		for begin := time.Now(); time.Since(begin) < rTime; {

			shareLock.Lock()
			time.Sleep(3 * time.Nanosecond)
			shareLock.Unlock()

			count++
		}

		fmt.Printf("greedy worker was able to execute %v work loops\n", count)
	}

	//只有在需要的时候才锁定
	politeWorker := func() {
		defer wg.Done()

		count := 0
		for begin := time.Now(); time.Since(begin) < rTime; {

			shareLock.Lock()
			time.Sleep(1 * time.Nanosecond)
			shareLock.Unlock()

			shareLock.Lock()
			time.Sleep(1 * time.Nanosecond)
			shareLock.Unlock()

			shareLock.Lock()
			time.Sleep(1 * time.Nanosecond)
			shareLock.Unlock()

			count++
		}

		fmt.Printf("Polite worker was able to execute %v work loops.\n", count)
	}

	wg.Add(2)
	// greedy不必要的扩大了对共享锁的控制，并且(通过饥饿)阻碍了polite有效的执行
	go greedyWorker() // greedy worker was able to execute 665383 work loops
	go politeWorker() // Polite worker was able to execute 354811 work loops.
	wg.Wait()
}
