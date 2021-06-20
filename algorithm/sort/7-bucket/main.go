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

// 桶排序：利用函数的映射关系，高效与否在于这个映射函数的确定
// 	算法动画： https://www.cs.usfca.edu/~galles/visualization/BucketSort.html
func bucketSort(nums []int, bucketSize int) []int {
	if len(nums) <= 1 {
		return nums
	}

	// 生成桶
	maxValue, minValue := nums[0], nums[0]
	for _, item := range nums {
		if item > maxValue {
			maxValue = item
		}
		if item < minValue {
			minValue = item
		}
	}
	bucketCount := (maxValue - minValue)/bucketSize + 1
	fmt.Printf("bucketCount = %d\n", bucketCount)
	bucketList := make([][]int, bucketCount)

	// 数据插入到bucket中
	for _, item := range nums {
		tmpArr := bucketList[(item - minValue)/bucketSize]
		if len(tmpArr) == 0 {
			bucketList[(item - minValue)/bucketSize] = append(tmpArr, item)
			continue
		}

		// 数组中有元素了, 则要在有序的元素中插入进去, 也就是实现一个插入排序
		var newSlice []int
		for key, sVal := range tmpArr {
			if (len(tmpArr)-1) == key && item > sVal {
				newSlice = append(tmpArr, item)
			}
			if item < sVal { //小于当前元素，则需要插入到这里了
				newSlice = append(tmpArr[:key], append([]int{item}, tmpArr[key+1:]...)...)
				continue
			}
		}
		bucketList[(item - minValue)/bucketSize] = newSlice

	}

	i := 0
	for _, bucketArr := range bucketList {
		for _, item := range bucketArr {
			nums[i] = item
			i++
		}
	}

	return nums
}

// 桶排序
func main() {
	nums := getRandArray(10)
	fmt.Println(nums)
	fmt.Println(bucketSort(nums, len(nums)))

}
