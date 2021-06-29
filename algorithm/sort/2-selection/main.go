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

// 选择排序: 每次标记最小的，然后与当前排序的索引交换，每次都是
func selectSort(nums []int) []int {
	l := len(nums)
	for i := 0; i < l; i++ {
		//core: 标记未排序元素的索引位置，最后方便用来做元素交换
		minIndex := i
		for j := i + 1; j < l; j++ {
			if nums[j] < nums[minIndex] {
				minIndex = j
			}
		}
		// swap minIndexVal
		nums[i], nums[minIndex] = nums[minIndex], nums[i]
	}
	return nums
}

func main() {
	nums := getRandArray(20)
	fmt.Println(nums)
	fmt.Println(selectSort(nums))
}
