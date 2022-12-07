package main

// p50. Pow(x, n)
// 实现 pow(x, n) ，即计算 x 的整数 n 次幂函数（即，xn ）。
//
// 
//
//示例 1：
//
//输入：x = 2.00000, n = 10
//输出：1024.00000
//示例 2：
//
//输入：x = 2.10000, n = 3
//输出：9.26100
//示例 3：
//
//输入：x = 2.00000, n = -2
//输出：0.25000
//解释：2-2 = 1/22 = 1/4 = 0.25
// 
//
//提示：
//
//-100.0 < x < 100.0
//-231 <= n <= 231-1
//-104 <= xn <= 104
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/powx-n
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {

}

// 优化递归代码
func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	if n < 0 {
		return myPow(x, -n)
	}

	if n % 2 == 1 {
		return x * myPow(x, n-1)
	}
	return myPow(x*x, n/2)
}


func myPowV2(x float64, n int) float64 {
	if n >= 0 {
		return helper(x, n)
	}
	return 1.0 / helper(x, -n)
}

// 折半递归
func helper(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	y := helper(x, n/2)
	if n % 2 == 0 {
		return y * y
	}
	return y * y * x
}
