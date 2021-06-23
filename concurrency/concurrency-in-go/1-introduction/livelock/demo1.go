package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// 活锁是主动执行并发操作的程序，但这些操作无法向前移动程序的状态
// 主要为了解释：加锁失败，放弃锁，再加锁
func main() {
	cadence := sync.NewCond(&sync.Mutex{})
	go func() {
		for range time.Tick(1 * time.Millisecond) {
			cadence.Broadcast()
		}
	}()

	takeStep := func() {
		cadence.L.Lock()
		cadence.Wait()
		cadence.L.Unlock()
	}

	tryDir := func(name string, dirName string, dir *int32) bool {
		fmt.Printf(name + " " + dirName + "\n")
		atomic.AddInt32(dir, 1) // <2>
		takeStep()              // <3>
		if atomic.LoadInt32(dir) == 1 {
			fmt.Printf("%v ---> success\n", name)
			return true
		}
		takeStep()
		atomic.AddInt32(dir, -1) // <4>
		return false
	}

	var left, right int32
	tryLeft := func(name string) bool { return tryDir(name, "left", &left) }
	tryRight := func(name string) bool { return tryDir(name, "right", &right) }

	walk := func(wg *sync.WaitGroup, name string) {
		defer wg.Done()
		fmt.Printf("%v is trying to scoot:", name)
		for i := 0; i < 5; i++ { // <1>
			if tryLeft(name) || tryRight(name) { // <2>
				return
			}
		}
		fmt.Printf("\n%v tosses her hands up in exasperation!", name)
	}

	var wg sync.WaitGroup // <3>
	wg.Add(2)
	go walk(&wg, "Alice")
	go walk(&wg, "Barbara")
	wg.Wait()

}
