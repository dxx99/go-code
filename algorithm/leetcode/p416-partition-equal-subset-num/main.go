package main

import "fmt"

// p416 分隔等和子集
// 给你一个 只包含正整数 的 非空 数组 nums 。请你判断是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。
//
// 
//
//示例 1：
//
//输入：nums = [1,5,11,5]
//输出：true
//解释：数组可以分割成 [1, 5, 5] 和 [11] 。
//示例 2：
//
//输入：nums = [1,2,3,5]
//输出：false
//解释：数组不能分割成两个元素和相等的子集。
// 
//
//提示：
//
//1 <= nums.length <= 200
//1 <= nums[i] <= 100
//
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/partition-equal-subset-sum
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	fmt.Println(solve01bag([]int{1,3,4}, []int{15,20,30}, 4))
	fmt.Println(solve01bagV2([]int{1,3,4}, []int{15,20,30}, 4))
	fmt.Println(solve01bagV3([]int{1,3,4}, []int{15,20,30}, 4))

}

// 先计算容量 sum/2
// dp[j] = max(dp[j], dp[j-num[i]] + nums[i])
// 初始化 0
func canPartition(nums []int) bool {
	
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	if sum % 2 != 0 {
		return false
	}

	// 背包容量
	target := sum/2

	//dp数组
	dp := make([]int, target+1)
	
	// 初始化，不需要
	
	//遍历
	for i := 0; i < len(nums); i++ {
		for j := target; j >= nums[i]; j-- {
			dp[j] = max(dp[j], dp[j-target]+nums[i])
		}
	}

	if dp[target] == target {
		return true
	}
	return false
}



// 01背包问题
// weight 对应背包的重量
// value 对应背包的价值
// 求最大重量maxWeight所对应的最大价值
func solve01bag(weight, value []int, maxWeight int) int {
	// dp[i][j] i个物品，j的重量，得到的最大价值
	// i表示物品的数量，j表示重量
	dp := make([][]int, len(weight))
	for i := range dp {
		dp[i] = make([]int, maxWeight+1)
	}

	// 初始化
	for i := maxWeight ; i >= weight[0]; i-- {
		dp[0][i] = dp[0][i - weight[0]] + value[0]
	}

	//递推公式, 当前元素 = 上一个元素，放不放当前重量的物品
	//dp[i][j] = max(dp[i-1][j], dp[i-1][j-weight[i]] + value[i])

	for i := 1; i < len(weight); i++ {
		for j := maxWeight; j >= weight[i] ; j-- {
			dp[i][j] = max(dp[i-1][j], dp[i-1][j - weight[i]] + value[i])
		}
	}

	fmt.Println(dp)
	return dp[len(weight)-1][maxWeight]
}

func solve01bagV2(weight, value []int, bagweight int) int {
	// 定义dp数组
	dp := make([][]int, len(weight))
	for i, _ := range dp {
		dp[i] = make([]int, bagweight+1)
	}
	// 初始化
	for j := bagweight; j >= weight[0]; j-- {
		dp[0][j] = dp[0][j-weight[0]] + value[0]
	}
	// 递推公式
	for i := 1; i < len(weight); i++ {
		//正序,也可以倒序
		for  j := weight[i];j<= bagweight ; j++ {
			dp[i][j] = max(dp[i-1][j], dp[i-1][j-weight[i]]+value[i])
		}
	}

	fmt.Println(dp)


	return dp[len(weight)-1][bagweight]
}

// 用滚动数组解决01背包问题
func solve01bagV3(weight, value []int, bagWeight int) int {
	// 定义dp数组 表示背包容量对应的价值
	dp := make([]int, bagWeight+1)

	// 初始化, 没有放物品时，价值为零
	dp[0] = 0

	// 递推公式
	// dp[j] = max(dp[j], dp[j-weight[i]] + value[i])
	for i := 0; i < len(weight); i++ {	// 先遍历物品
		for j := bagWeight; j >= weight[i] ; j-- {		// 这里不能覆盖
			dp[j] = max(dp[j], dp[j-weight[i]] + value[i])
		}

	}

	fmt.Println(dp)
	return dp[bagWeight]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}