package main

import (
	"fmt"
)
//给定一个数组 A[0,1,…,n-1]，请构建一个数组 B[0,1,…,n-1]，其中 B[i] 的值是数组 A 中除了下标 i 以外的元素的积, 即 B[i]=A[0]×A[1]×…×A[i-1]×A[i+1]×…×A[n-1]。不能使用除法。
//
// 
//
//示例:
//
//输入: [1,2,3,4,5]
//输出: [120,60,40,30,24]
// 
//
//提示：
//
//所有元素乘积之和不会溢出 32 位整数
//a.length <= 100000
//
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/gou-jian-cheng-ji-shu-zu-lcof
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	fmt.Println(constructArr([]int{}))
}

// 分别得到左下角和右上角
func constructArr(a []int) []int {
	if len(a) == 0 {
		return []int{}
	}
	leftArr, rightArr := make([]int, len(a)), make([]int, len(a))

	// 左上
	leftArr[0] = 1
	for i := 1; i < len(a); i++ {
		leftArr[i] = a[i-1]*leftArr[i-1]
	}
	// 右下
	rightArr[len(a)-1] = 1
	for i := len(a)-2; i >= 0; i-- {
		rightArr[i] = a[i+1]*rightArr[i+1]
	}

	// 再将左右相乘等到结果
	ans := make([]int, len(a))
	for i := 0; i < len(a); i++ {
		ans[i] = leftArr[i] * rightArr[i]
	}

	return ans
}

func constructArrV2(a []int) []int {
	total := 1
	zeroNum := 0
	lastZeroIndex := 0
	for i := 0; i < len(a); i++ {
		if a[i] == 0 {	// 对零进行特殊处理
			zeroNum++
			lastZeroIndex = i
			continue
		}
		total *= a[i]
	}
	ans := make([]int, len(a))
	if zeroNum > 1 {
		return ans
	}
	if zeroNum == 1 {
		ans[lastZeroIndex] = total
		return ans
	}

	for i := 0; i < len(a); i++ {
		ans[i] = total/a[i]
	}
	return ans
}

