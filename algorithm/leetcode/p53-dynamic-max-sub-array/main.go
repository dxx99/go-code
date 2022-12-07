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
	fmt.Println(maxSubArrayVdg(nums))
	fmt.Println(maxSubArrayVtx(nums))
}


// 分治算法
func maxSubArrayDg(nums []int) int {
	return get(nums, 0, len(nums)-1).mSum
}

type Status struct {
	lSum int	// 表示 [l,r] 内以 ll 为左端点的最大子段和
	rSum int	// 表示 [l,r] 内以 rr 为右端点的最大子段和
	mSum int	// 表示 [l,r] 内的最大子段和
	iSum int	// 表示 [l,r] 的区间和
}

func get(nums []int, l, r int) Status {
	if l == r {		// 只有一个元素
		return Status{nums[l], nums[l], nums[l], nums[l]}
	}

	mid := (l + r) >> 1
	lSub := get(nums, l, mid)
	rSub := get(nums, mid+1, r)
	return pushUp(lSub, rSub)
}

// 合并两个元素的
func pushUp(l, r Status) Status {
	iSum := l.iSum + r.iSum
	lSum := max(l.lSum, l.iSum+r.lSum)
	rSum := max(r.rSum, r.iSum + l.rSum)
	mSum := max(max(l.mSum, r.mSum), l.rSum+r.lSum)
	return Status{
		lSum: lSum,
		rSum: rSum,
		mSum: mSum,
		iSum: iSum,
	}
}











// 贪心算法
// 如果当前元素的和大于零，则加上当前元素之后的和最大值
// 如果当前元素的和小于零，则要丢弃当前元素的值
func maxSubArrayVtx(nums []int) int {
	sum, cur := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		if cur < 0 {
			cur = nums[i]
		}else {
			cur += nums[i]
		}

		if cur > sum {
			sum = cur
		}
	}
	return sum
}

func maxSubArrayVdg(nums []int) int {
	m := math.MinInt
	dp := make([]int, len(nums))

	// 当前元素减一的和大于零
	// dp[i] = max(dp[i-1], 0) + n

	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i-1], 0) + nums[i]
		// 存储最大值
		if dp[i] > m {
			m = dp[i]
		}
	}
	return m
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

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

// dynamicProgramming
// 如果前一个元素大于零，则将其加到当前元素上
// f(n) = max(f(n-1), 0) + n
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
