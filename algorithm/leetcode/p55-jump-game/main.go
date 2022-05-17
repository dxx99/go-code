package main

import "fmt"

// p55 跳跃游戏
// 给定一个非负整数数组 nums ，你最初位于数组的 第一个下标 。
//
//数组中的每个元素代表你在该位置可以跳跃的最大长度。
//
//判断你是否能够到达最后一个下标。
//
// 
//
//示例 1：
//
//输入：nums = [2,3,1,1,4]
//输出：true
//解释：可以先跳 1 步，从下标 0 到达下标 1, 然后再从下标 1 跳 3 步到达最后一个下标。
//示例 2：
//
//输入：nums = [3,2,1,0,4]
//输出：false
//解释：无论怎样，总会到达下标为 3 的位置。但该下标的最大跳跃长度是 0 ， 所以永远不可能到达最后一个下标。
// 
//
//提示：
//
//1 <= nums.length <= 3 * 104
//0 <= nums[i] <= 105
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/jump-game
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	fmt.Println(canJump([]int{5,9,3,2,1,0,2,3,3,1,0,0}))
	fmt.Println(canJump([]int{2,3,1,1,4}))
	fmt.Println(canJump([]int{3,2,1,0,4}))
	fmt.Println("v2-----------------")
	fmt.Println(canJumpV2([]int{5,9,3,2,1,0,2,3,3,1,0,0}))
	fmt.Println(canJumpV2([]int{2,3,1,1,4}))
	fmt.Println(canJumpV2([]int{3,2,1,0,4}))
}

func canJump(nums []int) bool {
	cover := 0
	if len(nums) == 0  {
		return true
	}

	// 当前元素能覆盖的范围
	for i := 0; i <= cover; i++ {
		if i + nums[i] > cover {
			cover = i + nums[i]
		}

		if cover >= len(nums)-1 {
			return true
		}
	}

	return false
}

// 动态规划求解
func canJumpV2(nums []int) bool {
	// dp 数组, 元素为i的值对应的是否能被覆盖
	dp := make([]bool, len(nums))

	// 初始化
	dp[0] = true

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {			//
			if dp[j] && j + nums[j] >= i {	// 通过j点，可以到达i点，说明是可以到达这个点的
				dp[i] = true
				break
			}
		}
	}

	return dp[len(nums)-1]
}
