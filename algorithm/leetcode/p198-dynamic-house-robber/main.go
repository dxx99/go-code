package main

import "fmt"

// 198. 打家劫舍
// 你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
//
//给定一个代表每个房屋存放金额的非负整数数组，计算你 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。
//
// 
//
//示例 1：
//
//输入：[1,2,3,1]
//输出：4
//解释：偷窃 1 号房屋 (金额 = 1) ，然后偷窃 3 号房屋 (金额 = 3)。
//     偷窃到的最高金额 = 1 + 3 = 4 。
//示例 2：
//
//输入：[2,7,9,3,1]
//输出：12
//解释：偷窃 1 号房屋 (金额 = 2), 偷窃 3 号房屋 (金额 = 9)，接着偷窃 5 号房屋 (金额 = 1)。
//     偷窃到的最高金额 = 2 + 9 + 1 = 12 。
// 
//
//提示：
//
//1 <= nums.length <= 100
//0 <= nums[i] <= 400
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/house-robber
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {

	fmt.Println([]int{1,2,3,1}[:4-1])

	fmt.Println(rob([]int{1,2,3,1}))
	fmt.Println(rob([]int{2,1}))
	fmt.Println(rob([]int{1,2}))

	fmt.Println("V2........................")

	fmt.Println(robV2([]int{1,2,3,1}))
	fmt.Println(robV2([]int{2,1}))
	fmt.Println(robV2([]int{1,2}))
	fmt.Println(robV2([]int{1}))
}

// dynamicProgramming：推出公式 []int{1,2,3,1}
// S表示偷多少间房的最大值，H表示当前房屋的财富
// 第1间房   S0 = H0 = 1
// 第2间房   S1 = Max(S0, H1) = 2
// 第3间房   S2 = Max(S1, S0 + H2) = 4
// 第4间房   S3 = Max(S2, S1 + H3) = 4
// 第N间房   Sn = Max(Sn-1, Sn-2 + Hn)
func rob(nums []int) int {
	l := len(nums)

	// 处理数组为空的结果
	if l == 0 {
		return 0
	}

	// 处理前面两种特殊的值
	s := make([]int, l)

	for i := 0; i < l; i++ {
		// 处理两种特殊的条件
		if i == 0 {
			s[i] = nums[i]
			continue
		}
		if i == 1 {
			if s[i-1] > nums[i] {
				s[i] = s[i-1]
			} else {
				s[i] = nums[i]
			}
			continue
		}

		// Sn = Max(Sn-1, Sn-2 + Hn)
		if s[i-2] + nums[i] > s[i-1] {
			s[i] = s[i-2] + nums[i]
		}else {
			s[i] = s[i-1]
		}
	}
	return s[l-1]
}

// dynamicProgramming + 滚动数组
func robV2(nums []int) int {
	// 前后两个值来替代数组
	first, second := 0, 0

	for i, num := range nums {
		if i == 0 {
			first, second = num, num	// 防止只有一个元素的求解问题
			continue
		}
		if i == 1 {
			if first > num {
				second = first
			} else {
				second = num
			}
			continue
		}

		if first + num > second {
			first, second = second, first + num
		}else {
			first = second
		}
	}

	return second
}

func robV3(nums []int) int {
	dp := make([]int, len(nums)+1)

	// 递推函数
	// dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	if len(nums) == 1 {
		return nums[0]
	}

	// 初始化
	dp[1], dp[2] = nums[0], max(nums[0],nums[1])

	for i := 3; i <= len(nums); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i-1])
	}

	return dp[len(nums)]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
