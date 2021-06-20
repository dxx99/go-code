package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 生成随机数组
func getRandArray(n int) []int {
	rand.Seed(time.Now().UnixNano())

	var res []int
	for i := 0; i < n; i++ {
		res = append(res, rand.Intn(1e3))
	}
	return res
}

// 希尔排序：插入排序的改进版本，它会优先比较距离较远的元素，希尔排序又叫缩小增量排序
// 关键字：增量，步长, 比较交换
// 算法描述： 先将整个待排序的记录序列分隔成为若干子序列分别进行直接插入排序
// 1. 选择一个增量序列t1，t2，…，tk，其中ti>tj，tk=1；
// 2. 按增量序列个数k，对序列进行k 趟排序；
// 3. 每趟排序，根据对应的增量ti，将待排序列分割成若干长度为m 的子序列，分别对各子表进行直接插入排序。仅增量因子为1 时，整个序列作为一个表来处理，表长度即为整个序列的长度。
func shellSort(nums []int) []int {
	inc := 2	// 增量
	step := len(nums)/inc  //初始步长

	for step >= 1 {
		for i := step; i < len(nums); i++ {
			for j := i; j >= step; j -= step {
				if nums[j] < nums[j - step] {
					// 右边比左边的值大，则交换
					nums[j], nums[j-step] = nums[j-step], nums[j]
				}
			}
		}
		// 对每轮步长的替换
		step = step/inc
	}
	return nums
}


// 希尔排序
// 动画： https://www.cs.usfca.edu/~galles/visualization/ComparisonSort.html
func main() {
	nums := getRandArray(11)
	fmt.Println(nums)
	fmt.Println(shellSort(nums))
}
