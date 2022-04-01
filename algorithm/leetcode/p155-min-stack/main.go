package main

import (
	"fmt"
	"math"
)

//155. 最小栈
// 设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。
//
//实现 MinStack 类:
//
//MinStack() 初始化堆栈对象。
//void push(int val) 将元素val推入堆栈。
//void pop() 删除堆栈顶部的元素。
//int top() 获取堆栈顶部的元素。
//int getMin() 获取堆栈中的最小元素。
// 
//
//示例 1:
//
//输入：
//["MinStack","push","push","push","getMin","pop","top","getMin"]
//[[],[-2],[0],[-3],[],[],[],[]]
//
//输出：
//[null,null,null,null,-3,null,0,-2]
//
//解释：
//MinStack minStack = new MinStack();
//minStack.push(-2);
//minStack.push(0);
//minStack.push(-3);
//minStack.getMin();   --> 返回 -3.
//minStack.pop();
//minStack.top();      --> 返回 0.
//minStack.getMin();   --> 返回 -2.
// 
//
//提示：
//
//-231 <= val <= 231 - 1
//pop、top 和 getMin 操作总是在 非空栈 上调用
//push, pop, top, and getMin最多被调用 3 * 104 次
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/min-stack
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	// ["MinStack","push","push","push","top","pop","getMin","pop","getMin","pop","push","top","getMin","push","top","getMin","pop","getMin"]
	//[[],[2147483646],[2147483646],[2147483647],[],[],[],[],[],[],[2147483647],[],[],[-2147483648],[],[],[],[]]
	minStack := Constructor()
	minStack.Push(2147483646)
	minStack.Push(2147483646)
	minStack.Push(2147483647)
	fmt.Println(minStack.Top())
	minStack.Pop()
	minStack.Pop()
	minStack.Pop()
}

// MinStack 也可以使用双栈，也就是一个辅助栈来实现这个
type MinStack struct {
	stack []int
	minVal int
}

func Constructor() MinStack {
	return MinStack{stack: make([]int, 0), minVal: math.MaxInt64}
}

// Push void push(int val) 将元素val推入堆栈
func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	if this.minVal > val {
		this.minVal = val
	}
}

// Pop void pop() 删除堆栈顶部的元素。
func (this *MinStack) Pop()  {
	length := len(this.stack)
	if length > 0 {
		// 删除栈顶部元素
		last := this.stack[length-1]
		this.stack = this.stack[:length-1]

		// 由于删除了最小值，则需要重新找一个最小值
		if this.minVal == last {
			if length - 1 > 0 {
				this.minVal = this.stack[0]
				for i := 1; i < length-1; i++ {
					if this.minVal > this.stack[i] {
						this.minVal = this.stack[i]
					}
				}
			}else{
				this.minVal = math.MaxInt64
			}
		}
	}

}


func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}


func (this *MinStack) GetMin() int {
	return this.minVal
}

