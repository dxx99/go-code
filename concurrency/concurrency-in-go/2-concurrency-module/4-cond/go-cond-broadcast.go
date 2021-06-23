package main

import (
	"fmt"
	"sync"
)

type Button struct {
	Clicked *sync.Cond
}

var button = Button{Clicked: sync.NewCond(&sync.Mutex{})}

func main() {
	subscribe := func(c *sync.Cond, fn func()) {
		var wg sync.WaitGroup // 这里我们创建一个WaitGroup。 这只是为了确保我们的程序在写入标准输出之前不会退出。
		wg.Add(1)
		go func() {
			wg.Done() // 这里不能使用defer wg.Done() 不然协程会deadlock 没法方式
			c.L.Lock()
			defer c.L.Unlock()
			c.Wait()
			fn()
		}()
		wg.Wait()
	}

	var waitG sync.WaitGroup
	waitG.Add(3)
	subscribe(button.Clicked, func() {
		waitG.Done()
		fmt.Println("Maximizing window.")
	})
	subscribe(button.Clicked, func() {
		waitG.Done()
		fmt.Println("Displaying annoying dialog box!")
	})
	subscribe(button.Clicked, func() {
		waitG.Done()
		fmt.Println("Mouse clicked.")
	})

	button.Clicked.Broadcast()	//唤醒因为条件阻塞的goroutine
	waitG.Wait()

}
