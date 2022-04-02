package main

import "fmt"

// p204 统计质数
// 质数的定义： 大于1的自然数，只能被自己或1整除，不能被其他的数整除的数
// 给定整数 n ，返回 所有小于非负整数 n 的质数的数量 。
//
// 
//
//示例 1：
//
//输入：n = 10
//输出：4
//解释：小于 10 的质数一共有 4 个, 它们是 2, 3, 5, 7 。
//示例 2：
//
//输入：n = 0
//输出：0
//示例 3：
//
//输入：n = 1
//输出：0
// 
//
//提示：
//
//
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/count-primes
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	fmt.Println(countPrimes(10))
	fmt.Println(countPrimes(0))
	fmt.Println(countPrimes(1))
}

func countPrimes(n int) (cnt int) {
	for x := 2; x < n; x++ {
		if isPrime(x) {
			cnt++
		}
	}
	return
}

// 质数判断的关键函数
func isPrime(x int) bool {
	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}