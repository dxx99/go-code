package main

// p213 打家劫舍II
// 你是一个专业的小偷，计划偷窃沿街的房屋，每间房内都藏有一定的现金。这个地方所有的房屋都 围成一圈 ，这意味着第一个房屋和最后一个房屋是紧挨着的。同时，相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警 。
//
//给定一个代表每个房屋存放金额的非负整数数组，计算你 在不触动警报装置的情况下 ，今晚能够偷窃到的最高金额。
//
// 
//
//示例 1：
//
//输入：nums = [2,3,2]
//输出：3
//解释：你不能先偷窃 1 号房屋（金额 = 2），然后偷窃 3 号房屋（金额 = 2）, 因为他们是相邻的。
//示例 2：
//
//输入：nums = [1,2,3,1]
//输出：4
//解释：你可以先偷窃 1 号房屋（金额 = 1），然后偷窃 3 号房屋（金额 = 3）。
//     偷窃到的最高金额 = 1 + 3 = 4 。
//示例 3：
//
//输入：nums = [1,2,3]
//输出：3
// 
//
//提示：
//
//1 <= nums.length <= 100
//0 <= nums[i] <= 1000
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/house-robber-ii
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {

}

// 打家劫舍成环问题，分开情况解决
func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums)==2 {
		return max(nums[0], nums[1])
	}

	res1 := robRange(nums, 0)
	res2 := robRange(nums, 1)
	return max(res1, res2)
}

func robRange(nums []int, start int) int {
	dp := make([]int, len(nums))
	dp[1] = nums[start]

	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-2]+nums[i-1+start], dp[i-1])
	}
	return dp[len(nums)-1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
