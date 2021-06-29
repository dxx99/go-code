package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// 通过字符串生成数组
func genArray() []int {
	str := "3,44,38,5,47,15,36,26,27,2,46,4,19,50,48"
	strArr := strings.Split(str, ",")

	var res []int
	for _, s := range strArr {
		n, _ := strconv.Atoi(s)
		res = append(res, n)
	}
	return res
}

// 生成随机数组
func getRandArray(n int) []int {
	rand.Seed(time.Now().UnixNano())

	var res []int
	for i := 0; i < n; i++ {
		res = append(res, rand.Intn(1e3))
	}
	return res
}

// 冒泡排序：每次左右两边元素进行比较(左大右就交换)，每轮确定把最后一个元素确定排序位置
// 01, 12, 23, 34, 56, 67, 78
// 01, 12, 23, 34, 56, 67
// 01, 12, 23, 34, 56
// 01, 12, 23, 34
// ...
func bubbleSort(nums []int) []int {
	l := len(nums)
	for i := 0; i < l; i++ {

		// Core: no-sort-data pre round, and Maximum sorting range
		for k := 0; k < l-1-i; k++ {
			// Compare Exchange
			if nums[k] > nums[k+1] {
				nums[k], nums[k+1] = nums[k+1], nums[k]
			}
		}
	}
	return nums
}

// Animation： https://visualgo.net/zh/sorting
func main() {
	aNums := getRandArray(10)
	fmt.Println(aNums)
	fmt.Println(bubbleSort(aNums))
}
