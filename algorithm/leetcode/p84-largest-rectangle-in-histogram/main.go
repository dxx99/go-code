package main

import "fmt"

//给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
//
//求在该柱状图中，能够勾勒出来的矩形的最大面积。
//
// 
//
//示例 1:
//
//
//
//输入：heights = [2,1,5,6,2,3]
//输出：10
//解释：最大的矩形为图中红色区域，面积为 10
//示例 2：
//
//
//
//输入： heights = [2,4]
//输出： 4
// 
//
//提示：
//
//1 <= heights.length <=105
//0 <= heights[i] <= 104
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/largest-rectangle-in-histogram
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	fmt.Println(largestRectangleArea([]int{1}))
	fmt.Println(largestRectangleArea([]int{1,1}))
	fmt.Println(largestRectangleArea([]int{2,1,5,6,2,3}))
	fmt.Println(largestRectangleAreaV2([]int{1}))
	fmt.Println(largestRectangleAreaV2([]int{1,1}))
	fmt.Println(largestRectangleAreaV2([]int{2,1,5,6,2,3}))
}

// dynamicProgramming
func largestRectangleAreaV2(heights []int) int {
	minLeft, minRight := make([]int, len(heights)), make([]int, len(heights))

	// 记录左右两边比当前柱子小的值
	minLeft[0] = -1
	for i := 1; i < len(heights); i++ {
		t := i-1
		for t >= 0 && heights[t] >= heights[i] {
			t = minLeft[t]
		}
		minLeft[i] = t
	}
	minRight[len(heights)-1] = len(heights)	// 注意这里初始化，防止下面for死循环
	for i := len(heights)-2; i >= 0; i-- {
		t := i + 1
		for t < len(heights) && heights[t] >= heights[i] {
			t = minRight[t]
		}
		minRight[i] = t
	}

	// 求和
	ans := 0
	for i := 0; i < len(heights); i++ {
		area := heights[i] * (minRight[i] - minLeft[i] - 1)
		if area > ans {
			ans = area
		}
	}
	return ans
}

// 单调栈【递减栈】
// 找到左右两边第一个小于当前柱子的位置
func largestRectangleArea(heights []int) int {
	// 往柱子前后加辅助柱子方便处理业务
	heights = append([]int{0}, heights...)
	heights = append(heights, 0)

	ans := 0
	stack := make([]int, 0)
	for i := 0; i < len(heights); i++ {
		for len(stack) != 0 && heights[i] < heights[stack[len(stack)-1]] {
			mid := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			// 处理业务
			if len(stack) != 0 {
				w := i - stack[len(stack)-1] - 1
				if w * heights[mid] > ans {
					ans = w * heights[mid]
				}
			}
		}

		stack = append(stack, i)
	}
	return ans
}
