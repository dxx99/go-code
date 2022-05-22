package main

import (
	"fmt"
	"math"
	"strconv"
)

// p738 单调递增的数字
// 当且仅当每个相邻位数上的数字 x 和 y 满足 x <= y 时，我们称这个整数是单调递增的。
//
//给定一个整数 n ，返回 小于或等于 n 的最大数字，且数字呈 单调递增 。
//
// 
//
//示例 1:
//
//输入: n = 10
//输出: 9
//示例 2:
//
//输入: n = 1234
//输出: 1234
//示例 3:
//
//输入: n = 332
//输出: 299
// 
//
//提示:
//
//0 <= n <= 109
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/monotone-increasing-digits
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	fmt.Println(monotoneIncreasingDigits(322211))
	fmt.Println(monotoneIncreasingDigits(100))
	fmt.Println(monotoneIncreasingDigits(10))
	fmt.Println(monotoneIncreasingDigits(1234))
	fmt.Println(monotoneIncreasingDigits(332))
}

// 贪心算法 优化代码
func monotoneIncreasingDigits(n int) int {
	bs := []byte(strconv.Itoa(n))

	if len(bs) <= 1 {
		return n
	}

	for i := len(bs)-1; i > 0; i-- {
		//这个是逆序， i-1是高位，i是低位，目的是让高位小于低位
		if bs[i-1] > bs[i] {
			bs[i-1]--
			for j := i ; j < len(bs); j++ {	//后面的全部置为9
				bs[j] = '9'
			}
		}
	}

	ans, _ := strconv.Atoi(string(bs))
	return ans
}

func monotoneIncreasingDigitsV2(n int) int {
	arr := make([]int, 0)
	for n >= 10 {
		arr = append(arr, n%10)
		n = n / 10
	}
	arr = append(arr, n)

	start := arr[0]
	for i := 1; i < len(arr); i++ {
		// 不需要处理
		if start >= arr[i] {
			start = arr[i]
			continue
		}

		// 需要做处理
		arr[i] -= 1
		cur := i
		for cur > 0 {
			arr[cur-1] = 9
			cur--
		}
		start = arr[i]
	}

	ans := arr[0]
	for i := 1; i < len(arr); i++ {
		ans += int(math.Pow10(i))*arr[i]
	}
	return ans
}