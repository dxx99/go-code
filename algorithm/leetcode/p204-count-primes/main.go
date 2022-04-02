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
	fmt.Println("V2................")
	fmt.Println(countPrimesV2(10))
	fmt.Println(countPrimesV2(0))
	fmt.Println(countPrimesV2(1))
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

// 质数判断
// 如果 X 是质数， 这 N*X 一定不是质数，则可以将这些元素排除
// 对于一个质数 X，如果按上文说的我们从 2X 开始标记其实是冗余的，应该直接从 X * X 开始标记，因为 2X,3X,… 这些数一定在 X 之前就被其他数的倍数标记过了，
// 例如 2 的所有倍数，3 的所有倍数 应该用这个数组记录，然后把这些元素全部除掉
// 时间复杂度 O(nloglog(n))
func countPrimesV2(n int) (cnt int) {
	// 先将所有的元素都设置为质数
	isPrimeArr := make([]bool, n)
	for i := range isPrimeArr {
		isPrimeArr[i] = true
	}

	// 下与n的质素数量
	for k := 2; k < n; k++ {
		if isPrimeArr[k] {
			cnt++

			// 关键，每次跳过当前元素 *2
			for j := 2*k; j < n; j += k {
				isPrimeArr[j] = false
			}
		}
	}
	return
}

// 线性刷
// 维护一个质数集合，从小到大遍历，如果当前的数X是质数，就加入到集合中
//
func countPrimesV3(n int) (cnt int) {
	primes := make([]int, 0)
	isPrimeArr := make([]bool, n)
	for i := range isPrimeArr {
		isPrimeArr[i] = true
	}

	for k := 2; k < n; k++ {
		if isPrimeArr[k] {
			primes = append(primes, k)
		}

		// 质数判断核心
		for _, p := range primes {
			if k * p >= n {
				break
			}
			isPrimeArr[k * p] = false
			if k % p == 0 {
				break
			}
		}
	}

	return len(primes)
}