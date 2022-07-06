package main

import "sort"

func main() {

}

// 300. 最长递增子序列
// https://leetcode.cn/problems/longest-increasing-subsequence/
// 分析：
//	- 当一个元素的时候，值为1
//  - 当二个元素的时候，比较nums(i) > nums[i-1], 则为2， 否则为1
//  - 当三个元素的时候，比较第三个元素与第二个元素，如果大于则+1， 在第二个元素的基础上
//	- 当第N个元素的时候，继续比较替换。。。
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// dp数组的定义：
	// 1. dp[i]表示以nums[i]这个数结尾的最长递增子序列的长度
	// 2. dp[i]如何求得，表示i前面 (0<j<i), nums[i] > nums[j]，表示前面j小于当前i的值+1
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}

	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j]  {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	ans := dp[0]
	for i := 1; i < len(dp); i++ {
		if dp[i] > ans {
			ans = dp[i]
		}
	}
	return ans
}
// 优化：二分查找 O(N*Log(N))
func lengthOfLISV2(nums []int) int {
	top := make([]int, len(nums))	// 存放堆的位置，只存放一个元素在堆上
	piles := 0 // 初始化牌堆的数量

	for i := 0; i < len(nums); i++ {
		poker := nums[i] 	// 当前要处理的元素值

		x := sort.SearchInts(top[:piles], poker)

		// 没有找到合适的牌堆，需要重新建一个
		if x == piles {
			piles++
		}
		top[x] = poker	// 只存储最上次的元素
	}
	return piles
}

