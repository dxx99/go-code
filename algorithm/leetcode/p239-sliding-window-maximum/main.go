package main

import (
	"container/list"
	"fmt"
)

// 239. 滑动窗口最大值
// 给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。
//你只可以看到在滑动窗口内的 k 个数字。
//滑动窗口每次只向右移动一位。
//
//返回 滑动窗口中的最大值 。
//
// 
//
//示例 1：
//
//输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
//输出：[3,3,5,5,6,7]
//解释：
//滑动窗口的位置                最大值
//---------------               -----
//[1  3  -1] -3  5  3  6  7       3
// 1 [3  -1  -3] 5  3  6  7       3
// 1  3 [-1  -3  5] 3  6  7       5
// 1  3  -1 [-3  5  3] 6  7       5
// 1  3  -1  -3 [5  3  6] 7       6
// 1  3  -1  -3  5 [3  6  7]      7
//示例 2：
//
//输入：nums = [1], k = 1
//输出：[1]
// 
//
//提示：
//
//1 <= nums.length <= 105
//-104 <= nums[i] <= 104
//1 <= k <= nums.length
//
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/sliding-window-maximum
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	fmt.Println(maxSlidingWindowV1([]int{1,3,-1,-3,5,3,6,7}, 3))
	fmt.Println(maxSlidingWindow([]int{1,3,-1,-3,5,3,6,7}, 3))
}

func maxSlidingWindow(nums []int, k int) []int {
	ans := make([]int, 0)
	maxQ := make([]int, 0)

	push := func(v int) {
		for len(maxQ) > 0 && maxQ[len(maxQ)-1] < v {
			maxQ = maxQ[:len(maxQ)-1]
		}
		maxQ = append(maxQ, v)
	}

	pop := func(v int) {
		if len(maxQ) > 0 && maxQ[0] == v {
			maxQ = maxQ[1:]
		}
	}

	// 初始化
	for i := 0; i < k-1; i++ {
		push(nums[i])
	}

	// 处理第k个元素

	for i := k-1; i < len(nums); i++ {
		push(nums[i])
		ans = append(ans, maxQ[0])
		pop(nums[i-k+1])
	}

	return ans
}

func maxSlidingWindowV1(nums []int, k int) []int {
	ans := make([]int, 0)
	window := NewMonotonicQueue()
	for i, num := range nums {
		// 先把前k-1个元素加入到单调(单调递减，最大的元素放到最前面)队列中
		if i < k-1 {
			window.Push(num)
			continue
		}

		// 加入第k个元素之后就需要存储单调队列的最大值
		window.Push(num)
		ans = append(ans, window.Max())

		// 移除左边的元素
		window.Pop(nums[i-k+1])
	}

	return ans
}

// MonotonicQueueInterface 单调队列api，保持一个单调递减的顺序
type MonotonicQueueInterface interface {
	Push(x int)	// 在队尾添加添加元素x【后面添加元素，要把前面比自己小的元素都删掉】
	Pop(x int)	// 队头元素如果是x, 删除
	Max() int	// 返回当前队列中的最大值
}

type MonotonicQueue struct {
	l *list.List
}

func NewMonotonicQueue() *MonotonicQueue {
	return &MonotonicQueue{list.New()}
}

func (m *MonotonicQueue) Push(x int) {
	for m.l.Len() > 0 && m.l.Back().Value.(int) < x {
		m.l.Remove(m.l.Back())
	}
	m.l.PushBack(x)
}

func (m *MonotonicQueue) Pop(x int) {
	if m.l.Len() > 0 && m.l.Front().Value.(int) == x {
		m.l.Remove(m.l.Front())
	}
}

func (m MonotonicQueue) Max() int {
	return m.l.Front().Value.(int)
}











