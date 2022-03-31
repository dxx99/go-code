package main

import (
	"fmt"
	"math"
)

// 53. 最大子数组和
// 给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
//
//子数组 是数组中的一个连续部分。
//
// 
//
//示例 1：
//
//输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
//输出：6
//解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。
//示例 2：
//
//输入：nums = [1]
//输出：1
//示例 3：
//
//输入：nums = [5,4,-1,7,8]
//输出：23
// 
//
//提示：
//
//1 <= nums.length <= 105
//-104 <= nums[i] <= 104
//
// 进阶：如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的 分治法 求解。
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/maximum-subarray
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	nums := []int{-2,1,-3,4,-1,2,1,-5,4}
	fmt.Println(maxSubArray(nums))
	fmt.Println(maxSubArrayV3(nums))
}

//TODO: 实现这个方法
// 两次循环, 每一轮中找到最大的
func maxSubArray(nums []int) int {
	l := len(nums)
	max := math.MinInt64
	for k := 0; k < l; k++ {
		sum := 0
		for j := k; j < l; j++ {
			sum += nums[j]
			if sum > max {
				max = sum
			}
		}
	}
	return max
}

// 动态规划
// 如果前一个元素大于零，则将其加到当前元素上
func maxSubArrayV2(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	max := nums[0]
	for k := 1; k < l; k++ {
		// 重置当前元素，这样就可以得出最大值的一堆数组了
		if nums[k] + nums[k-1] > nums[k] {
			nums[k] = nums[k] + nums[k-1]
		}
		if nums[k] > max {
			max = nums[k]
		}
	}

	return max
}

// 贪心算法
// 如果当前指针所指元素的和小于0，则丢弃当前元素之前的数列
func maxSubArrayV3(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	max, current := nums[0], nums[0]

	for k := 1; k < l; k++ {
		if current < 0 {
			current = 0
		}
		current += nums[k]
		if current > max {
			max = current
		}
	}

	return max
}
