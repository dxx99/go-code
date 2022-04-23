package main

import "fmt"

// 70. 爬楼梯
// 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
//
//每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
//
// 
//
//示例 1：
//
//输入：n = 2
//输出：2
//解释：有两种方法可以爬到楼顶。
//1. 1 阶 + 1 阶
//2. 2 阶
//示例 2：
//
//输入：n = 3
//输出：3
//解释：有三种方法可以爬到楼顶。
//1. 1 阶 + 1 阶 + 1 阶
//2. 1 阶 + 2 阶
//3. 2 阶 + 1 阶
// 
//
//提示：
//
//1 <= n <= 45
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/climbing-stairs
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	fmt.Println(climbStairsV2(2))
	fmt.Println(climbStairsV2(3))
	fmt.Println(climbStairsV2(100))
	fmt.Println(climbStairsV5(2))
	fmt.Println(climbStairsV5(3))
	fmt.Println(climbStairsV5(100))
}

// 递归请求
// 这个就是类似于斐波拉契队列
// f(n) = f(n-1) + f(n-2)  假设当前在第n个台阶，则他上一个台阶在n-1或者n-2
func climbStairs(n int) int {
	// 临界条件
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	return climbStairs(n-1) + climbStairs(n-2)
}

// 递归优化
func climbStairsV2(n int) int {
	// 临界条件
	m := make(map[int]int)
	return helper(n, m)
}

func helper(n int, m map[int]int) int {
	if v, ok := m[n]; ok {
		return v
	}
	if n == 1 {
		m[1] = 1
	}else if n == 2 {
		m[2] = 2
	}else {
		m[n] = helper(n-1, m) + helper(n-2, m)
	}
	return m[n]
}

// 使用迭代求解
func climbStairsV3(n int) int {
	if n == 1 {
		return 1
	}

	// 可以创建一个数组，记录各个结果的值
	f1, f2 := 1, 2
	for i := 3; i <= n; i++ {
		f1, f2 = f2, f1+f2
	}
	return f2
}

// 动态规划求解
func climbStairsV4(n int) int {
	dp := make([]int, n+1)		// 1. 定义dp数组

	dp[1], dp[2] = 1, 2		// 3. 初始化递推数组
	for i := 3; i <= n; i++ {	//4. 遍历，从什么时候开始，
		dp[i] = dp[i-1] + dp[i-2]	//2. 推到递推公司
	}
	return dp[n]
}

// 动态规划 + 完全背包求解
func climbStairsV5(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1

	arr := []int{1,2}
	for j := 0; j <= n; j++ {
		for i := 0; i < len(arr); i++ {
			if j >= arr[i] {
				dp[j] += dp[j - arr[i]]
			}
		}
	}

	return dp[n]
}
