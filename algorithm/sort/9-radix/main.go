package main

import (
	"fmt"
	"math"
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

// 基数排序: 按照低位先排序，然后收集，再按照高位排序，再收集，依次类推..
// 动画： https://www.cs.usfca.edu/~galles/visualization/RadixSort.html
// 算法描述：
// 1. 取得数组中的最大数，并取得位数
// 2. 从低位开始，每个位组成一个radix数组
// 3. 对radix数组进行计算
func radixSort(nums []int) []int {

	// 得到元素最大的位数
	maxValue := nums[0]
	for _, item := range nums {
		if item > maxValue {
			maxValue = item
		}
	}
	maxDigit := 1
	for float64(maxValue) > math.Pow10(maxDigit) {
		maxDigit++
	}

	// 准备一个数组存储radix数据

	mod, dev := 10, 1
	for i := 0; i < maxDigit; i++ {
		radixList := make([][]int, 10)

		//把数据放到radixList中
		for _, item := range nums {
			rKey := (item%mod) / dev
			radixList[rKey] = append(radixList[rKey], item)
		}

		k := 0
		for _, radixArr := range radixList {
			for _, item := range radixArr {
				nums[k] = item
				k++
			}
		}

		mod *= 10
		dev *= 10
	}
	return nums
}

func main() {
	nums := getRandArray(10)
	fmt.Println(nums)
	fmt.Println(radixSort(nums))
}
