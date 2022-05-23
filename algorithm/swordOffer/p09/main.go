package main

import "fmt"

func main() {

	obj := Constructor()
	obj.AppendTail(3)
	fmt.Println(obj.DeleteHead())
	fmt.Println(obj.DeleteHead())

}

// CQueue 使用双栈实现队列
type CQueue struct {
	pushStack []int
	popStack []int
}


func Constructor() CQueue {
	return CQueue{
		pushStack: make([]int, 0),
		popStack:  make([]int, 0),
	}
}

// AppendTail 向队列中添加元素
func (c *CQueue) AppendTail(value int)  {
	c.pushStack = append(c.pushStack, value)
}


// DeleteHead 从队列中头部删除元素，如果成功，则返回当前元素，失败返回-1
func (c *CQueue) DeleteHead() int {
	if len(c.popStack) > 0 {
		ans := c.popStack[len(c.popStack)-1]
		c.popStack = c.popStack[:len(c.popStack)-1]
		return ans
	}

	if len(c.pushStack) > 0 {
		ans := c.pushStack[0]
		for i := len(c.pushStack)-1; i > 0; i-- {
			c.popStack = append(c.popStack, c.pushStack[i])
		}
		c.pushStack = make([]int, 0)

		return ans
	}

	return -1
}


/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */