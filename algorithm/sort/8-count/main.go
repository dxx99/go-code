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

func countSort(nums []int) []int {
	maxValue := nums[0]
	for _, item := range nums {
		if item > maxValue {
			maxValue = item
		}
	}

	countList := make([]int, maxValue+1)

	for _, item := range nums {
		countList[item] += 1
	}

	i := 0
	for k, val := range countList {
		for val > 0 {
			nums[i] = k
			i++
			val--
		}
	}
	return nums
}

// 计数排序：核心是将输入的数据转化成键存储在额外开辟的数组空间中
// 	  要求：计数排序要求输入的数据必须有确定范围的整数
// 动画演示：https://www.cs.usfca.edu/~galles/visualization/CountingSort.html
func main() {
	nums := getRandArray(10)
	fmt.Println(nums)
	fmt.Println(countSort(nums))
}
