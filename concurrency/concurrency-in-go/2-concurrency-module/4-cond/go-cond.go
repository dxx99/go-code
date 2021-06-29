package main

import (
	"fmt"
	"sync"
	"time"
)

// 实现了一个条件变量，用于等待或宣布事件发生时goroutine的交汇点
// 事件：指两个或者多个goroutine之间的任何信号，仅指事件发生了，不包含任何其他信息
/*
通过死循环挂起的形式来实现，收到某个信号来进行等待，消耗cpu, 加入time.sleep()也影响性能
for conditionTrue() == false {

}
*/
func main() {
	cond := sync.NewCond(&sync.Mutex{})
	queue := make([]interface{}, 0, 10)

	removeFromQueue := func(delay time.Duration) {
		time.Sleep(delay)
		cond.L.Lock()
		x := queue[0]
		queue = queue[1:]
		fmt.Printf("remove item=%v from queue.\n", x)
		cond.L.Unlock()
		cond.Signal() // 发送信号，唤醒等待的goroutine
	}

	for i := 0; i < 10; i++ {
		cond.L.Lock()
		// 在这里我们检查队列的长度，以确认什么时候需要等待。由于removeFromQueue是异步的，for不满足时才会跳出，而if做不到重复判断，这一点很重要
		if len(queue) == 2 {
			cond.Wait() //5 调用wait, 这将阻塞main goroutine, 直到接受到信号
		}
		item := struct{ Id int }{Id: i}
		queue = append(queue, item)
		//fmt.Printf("add item=%v into queue\n", item)
		cond.L.Unlock()
		go removeFromQueue(1 * time.Second) //6 创建一个新的goroutine, 在1s之后将元素移除队列

	}

	time.Sleep(3 * time.Second)
}
