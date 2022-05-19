package main

import "fmt"

// p56 跳跃游戏II
// 给你一个非负整数数组 nums ，你最初位于数组的第一个位置。
//
//数组中的每个元素代表你在该位置可以跳跃的最大长度。
//
//你的目标是使用最少的跳跃次数到达数组的最后一个位置。
//
//假设你总是可以到达数组的最后一个位置。
//
// 
//
//示例 1:
//
//输入: nums = [2,3,1,1,4]
//输出: 2
//解释: 跳到最后一个位置的最小跳跃数是 2。
//     从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。
//示例 2:
//
//输入: nums = [2,3,0,1,4]
//输出: 2
// 
//
//提示:
//
//1 <= nums.length <= 104
//0 <= nums[i] <= 1000
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/jump-game-ii
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	fmt.Println(jump([]int{2,3,1,1,4}))
	fmt.Println(jump([]int{2,3,0,1,4}))

	fmt.Println("v2-------------")
	fmt.Println(jumpV2([]int{2,3,1,1,4}))
	fmt.Println(jumpV2([]int{2,3,0,1,4}))
}

// 动态规划求解
func jumpV2(nums []int) int {
	// dp数组
	dp := make([]int, len(nums))
	dp[0] = 0

	// 遍历
	index := 0
	for i := 1; i < len(nums); i++ {
		for i > index + nums[index] {   // 到达索引i点时，需要index最小值是多少
			index++
		}

		dp[i] = dp[index]+1
	}

	return dp[len(nums)-1]
}

// 贪心算法
// 看最大覆盖的范围：什么时候步数才一定要加一
// 思路：当前可移动的距离尽可能多走，如果还没到达终点，步数加一，整体：一步尽可能多走，从而达到最小步数
// 		要从覆盖范围出发，不管怎么跳，覆盖的范围一定是可以跳到的，以最小的步数增加覆盖的范围，覆盖的范围一段覆盖了终点，就得到最小步数
// 注意：这里要统计两个覆盖范围，当前一步的最大覆盖范围，下一步的最大覆盖
func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	ans := 0		// 记录走的最大步数

	curDist := 0	// 当前覆盖最远距离下标
	nextDist := 0 	// 下一步覆盖的最远距离下标
	for i := 0; i < len(nums)-1; i++ {			// 这里的len(nums)-1是关键
		nextDist = max(nums[i] + i, nextDist)	// 更新下一步覆盖的最远距离
		if i == curDist {
			curDist = nextDist
			ans++
		}
	}

	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
