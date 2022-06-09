package main

import (
	"fmt"
)

//把n个骰子扔在地上，所有骰子朝上一面的点数之和为s。输入n，打印出s的所有可能的值出现的概率。
//
// 
//
//你需要用一个浮点数数组返回答案，其中第 i 个元素代表这 n 个骰子所能掷出的点数集合中第 i 小的那个的概率。
//
// 
//
//示例 1:
//
//输入: 1
//输出: [0.16667,0.16667,0.16667,0.16667,0.16667,0.16667]
//示例 2:
//
//输入: 2
//输出: [0.02778,0.05556,0.08333,0.11111,0.13889,0.16667,0.13889,0.11111,0.08333,0.05556,0.02778]
// 
//
//限制：
//
//1 <= n <= 11
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/nge-tou-zi-de-dian-shu-lcof
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {

	//fmt.Println(dicesProbability(1))
	fmt.Println(dicesProbability(2))
}

func dicesProbability(n int) []float64 {
	// 初始化一颗股子
	dp := make([]float64, 6)
	for i := 0; i < len(dp); i++ {
		dp[i] = 1.0/6.0
	}

	for i := 2; i <= n; i++ {
		tmp := make([]float64, 5*i+1)
		for j := 0; j < len(dp); j++ {
			for k := 0; k < 6; k++ {
				tmp[j+k] += dp[j]/6.0	// 递推公式
			}
		}
		dp = tmp
	}

	return dp
}

func dicesProbabilityV2(n int) []float64 {

	min, max := n * 1, n*6
	backRes := make([]int, max-min+1)

	totalNum := 0
	nums := []int{1,2,3,4,5,6}
	track := make([]int, 0)
	var backtracking func()
	backtracking = func() {
		if len(track) == n {
			totalNum++
			backRes[sum(track) - min]++
			return
		}

		for i := 0; i < len(nums); i++ {
			track = append(track, nums[i])
			backtracking()

			track = track[:len(track)-1]
		}
	}

	backtracking()
	ans := make([]float64, max-min+1)
	for i := 0; i < len(ans); i++ {
		ans[i] = float64(backRes[i])/float64(totalNum)
	}
	return ans

}

func sum(nums []int) int {
	ans := 0
	for i := 0; i < len(nums); i++ {
		ans += nums[i]
	}
	return ans
}
