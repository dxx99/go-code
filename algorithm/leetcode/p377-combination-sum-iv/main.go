package main

import "fmt"

// p377 组合总和IV
// 给你一个由 不同 整数组成的数组 nums ，和一个目标整数 target 。请你从 nums 中找出并返回总和为 target 的元素组合的个数。
//
//题目数据保证答案符合 32 位整数范围。
//
// 
//
//示例 1：
//
//输入：nums = [1,2,3], target = 4
//输出：7
//解释：
//所有可能的组合为：
//(1, 1, 1, 1)
//(1, 1, 2)
//(1, 2, 1)
//(1, 3)
//(2, 1, 1)
//(2, 2)
//(3, 1)
//请注意，顺序不同的序列被视作不同的组合。
//示例 2：
//
//输入：nums = [9], target = 3
//输出：0
// 
//
//提示：
//
//1 <= nums.length <= 200
//1 <= nums[i] <= 1000
//nums 中的所有元素 互不相同
//1 <= target <= 1000
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/combination-sum-iv
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	fmt.Println(combinationSum4([]int{1,2,3}, 4))
}

func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)

	// 递推公式
	// dp[j] += dp[j-nums[i]]

	// 初始化, 总和为0时，也就是空数组组合
	dp[0] = 1

	// 因为有序, 需要先遍历背包，再遍历物品
	for j := 0; j <= target; j++ {
		for i := 0; i < len(nums); i++ {
			if j >= nums[i] {
				dp[j] += dp[j-nums[i]]
				//fmt.Println(dp)
			}
		}
	}

	return dp[target]
}
