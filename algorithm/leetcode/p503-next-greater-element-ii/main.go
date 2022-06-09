package main

import "fmt"

// 给定一个循环数组 nums （ nums[nums.length - 1] 的下一个元素是 nums[0] ），返回 nums 中每个元素的 下一个更大元素 。
//
//数字 x 的 下一个更大的元素 是按数组遍历顺序，这个数字之后的第一个比它更大的数，这意味着你应该循环地搜索它的下一个更大的数。如果不存在，则输出 -1 。
//
// 
//
//示例 1:
//
//输入: nums = [1,2,1]
//输出: [2,-1,2]
//解释: 第一个 1 的下一个更大的数是 2；
//数字 2 找不到下一个更大的数；
//第二个 1 的下一个最大的数需要循环搜索，结果也是 2。
//示例 2:
//
//输入: nums = [1,2,3,4,3]
//输出: [2,3,4,-1,4]
// 
//
//提示:
//
//1 <= nums.length <= 104
//-109 <= nums[i] <= 109
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/next-greater-element-ii
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	//fmt.Println(nextGreaterElements([]int{1,2,1}))
	fmt.Println(nextGreaterElements([]int{1,2,3,4,3}))
}


func nextGreaterElements(nums []int) []int {
	stack := make([]int, 0)
	nl := len(nums)
	ans := make([]int, nl)
	for i := 0; i < len(ans); i++ {
		ans[i] = -1
	}

	for i := 0; i < 2*nl; i++ {
		k := i % nl
		for len(stack) != 0 && nums[stack[len(stack)-1]] < nums[k] {
			top := stack[len(stack)-1]
			ans[top] = nums[k]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, k)
	}

	return ans
}