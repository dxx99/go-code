package main

import "fmt"

// 请定义一个队列并实现函数 max_value 得到队列里的最大值，要求函数max_value、push_back 和 pop_front 的均摊时间复杂度都是O(1)。
//
//若队列为空，pop_front 和 max_value 需要返回 -1
//
//示例 1：
//
//输入:
//["MaxQueue","push_back","push_back","max_value","pop_front","max_value"]
//[[],[1],[2],[],[],[]]
//输出: [null,null,null,2,1,2]
//示例 2：
//
//输入:
//["MaxQueue","pop_front","max_value"]
//[[],[],[]]
//输出: [null,-1,-1]
// 
//
//限制：
//
//1 <= push_back,pop_front,max_value的总操作数 <= 10000
//1 <= value <= 10^5
//
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/dui-lie-de-zui-da-zhi-lcof
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	obj := Constructor()
	obj.Push_back(1)
	obj.Push_back(2)
	fmt.Println(obj.Max_value())
	fmt.Println(obj.Pop_front())
	fmt.Println(obj.Max_value())
}

// MaxQueue 队列+单调双端队列
type MaxQueue struct {
	queue []int
	dq []int	// 单调双端递增队列
}


func Constructor() MaxQueue {
	return MaxQueue{make([]int, 0), make([]int, 0)}
}


func (m *MaxQueue) Max_value() int {
	if len(m.dq) == 0 {
		return -1
	}
	return m.dq[0]
}


func (m *MaxQueue) Push_back(value int)  {
	m.queue = append(m.queue, value)

	for len(m.dq) != 0 && m.dq[len(m.dq)-1] < value {
		m.dq = m.dq[:len(m.dq)-1]
	}
	m.dq = append(m.dq, value)
}


func (m *MaxQueue) Pop_front() int {
	if len(m.queue) == 0 {
		return -1
	}
	v := m.queue[0]
	m.queue = m.queue[1:]

	// 处理队列【队头元素】
	if len(m.dq) != 0 && m.dq[0] == v {
		m.dq = m.dq[1:]
	}

	return v
}


/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */
