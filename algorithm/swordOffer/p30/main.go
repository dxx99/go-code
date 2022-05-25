package main

import "math"

// 定义栈的数据结构，请在该类型中实现一个能够得到栈的最小元素的 min 函数在该栈中，调用 min、push 及 pop 的时间复杂度都是 O(1)。
//
// 
//
//示例:
//
//MinStack minStack = new MinStack();
//minStack.push(-2);
//minStack.push(0);
//minStack.push(-3);
//minStack.min();   --> 返回 -3.
//minStack.pop();
//minStack.top();      --> 返回 0.
//minStack.min();   --> 返回 -2.
// 
//
//提示：
//
//各函数的调用总次数不超过 20000 次
// 
//
//注意：本题与主站 155 题相同：https://leetcode-cn.com/problems/min-stack/
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/bao-han-minhan-shu-de-zhan-lcof
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {

}

// MinStack 使用双栈处理数据
type MinStack struct {
	bucket []int
	minArr    []int
}

func Constructor() MinStack {
	return MinStack{
		bucket: make([]int, 0),
		minArr:    make([]int, 0),
	}
}


func (e *MinStack) Push(x int)  {
	e.bucket = append(e.bucket, x)
	if len(e.minArr) == 0 {
		e.minArr = append(e.minArr, x)
	}else {
		cur := e.minArr[len(e.minArr)-1]
		if x <= cur {
			e.minArr = append(e.minArr, x)
		}
	}
}


func (e *MinStack) Pop()  {
	if len(e.bucket) == 0 {
		return
	}
	x := e.bucket[len(e.bucket)-1]
	if len(e.minArr) != 0 {
		m := e.minArr[len(e.minArr)-1]
		if m == x {
			e.minArr = e.minArr[:len(e.minArr)-1]
		}
	}
	e.bucket = e.bucket[:len(e.bucket)-1]
}


func (e *MinStack) Top() int {
	return e.bucket[len(e.bucket)-1]
}


func (e *MinStack) Min() int {
	if len(e.minArr) == 0 {
		return math.MaxInt64
	}
	return e.minArr[len(e.minArr)-1]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */