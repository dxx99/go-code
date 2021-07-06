package main

import (
	"errors"
	"fmt"
	"sync"
	"unsafe"
)

type queue interface {
	IsEmpty() bool
	Len() int
	Push(e Element) error	// 入队
	Pop() (Element, error)	// 出队
	Clear() error   		// 清空队列
	PrintStack()			// 打印队列
}

type cycleQueue struct {
	size     uint            // 队列大小
	dataSize uint            // 数据域的大小
	inIndex     uint            // 插入元素的索引值
	outIndex     uint            // 拿元素的索引值
	buf      unsafe.Pointer // 指向队列的数据域
	lock sync.Mutex			//
}

func (c cycleQueue) IsEmpty() bool {
	if c.dataSize == 0 {
		return true
	}
	return false
}

func (c *cycleQueue) Len() int {
	return int(c.dataSize)
}

func (c *cycleQueue) Push(e Element) error {
	if c.size == c.dataSize {
		return errors.New("queue is empty, push element failed")
	}
	c.lock.Lock()
	defer c.lock.Unlock()
	k := c.inIndex % c.size
	s :=  *(*[]Element)(c.buf)
	s[k] = e

	c.inIndex++
	c.dataSize++
	return nil
}

func (c *cycleQueue) Pop() (Element, error) {
	if c.dataSize <= 0 {
		return nil, errors.New("queue is empty, pop element failed")
	}
	c.lock.Lock()
	defer c.lock.Unlock()

	k := c.outIndex % c.size
	s :=  *(*[]Element)(c.buf)
	e := s[k]
	c.outIndex++
	c.dataSize--

	return e, nil
}

func (c *cycleQueue) Clear() error {
	if c.dataSize == 0 {
		return nil
	}
	c.lock.Lock()
	defer c.lock.Unlock()

	buf := make([]Element, c.size)
	c.buf = unsafe.Pointer(&buf)
	c.dataSize = 0
	c.inIndex = 0
	c.outIndex = 0
	return nil
}

func (c *cycleQueue) PrintStack() {
	s :=  *(*[]Element)(c.buf)
	fmt.Printf("size=%d, datasize=%d, inIndex=%d, outIndex=%d, buf=%v\n", c.size, c.dataSize, c.inIndex, c.outIndex, s)
}

type Element interface{}

func NewCycleQueue(size uint) *cycleQueue {
	buf := make([]Element, size)
	return &cycleQueue{
		size:     size,
		dataSize: 0,
		inIndex:  0,
		outIndex: 0,
		buf:      unsafe.Pointer(&buf),
		lock:     sync.Mutex{},
	}
}




// 环形队列
func main() {
	q := NewCycleQueue(3)
	fmt.Println(q.Push(1))
	fmt.Println(q.Push(2))
	fmt.Println(q.Push(3))
	q.PrintStack()	//size=3, datasize=3, inIndex=3, outIndex=0, buf=[1 2 3]

	fmt.Println(q.Push(4))	// queue is empty, push element failed
	fmt.Println(q.Pop())	// 1 <nil>
	q.PrintStack()	//size=3, datasize=2, inIndex=3, outIndex=1, buf=[1 2 3]
	fmt.Println(q.Push(4))
	q.PrintStack()	// size=3, datasize=3, inIndex=4, outIndex=1, buf=[4 2 3]
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())	//<nil> queue is empty, pop element failed
	fmt.Println(q.Push(5))
	fmt.Println(q.Push(6))
	fmt.Println(q.Push(7))
	q.PrintStack()	// size=3, datasize=3, inIndex=7, outIndex=4, buf=[7 5 6]

}
