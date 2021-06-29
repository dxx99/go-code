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

// 归并排序：是建立在归并操作的一种有效排序算法，采用的是分治法，将已有序的子序列合并，得到完全有序的序列
// 关键字： 分组，合并
// 1. 把长度为n的输入序列分成两个长度为n/2的子序列；
// 2. 对这两个子序列分别采用归并排序；
// 3. 将两个排序好的子序列合并成一个最终的排序序列
func mergeSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	mid := len(nums) / 2

	return func(left, right []int) []int {
		res := make([]int, 0)

		// 合并左右两边的数据
		for len(left) > 0 && len(right) > 0 {
			if left[0] > right[0] {
				res = append(res, right[0])
				right = right[1:]
			} else {
				res = append(res, left[0])
				left = left[1:]
			}
		}

		// 左右元素还有的情况
		if len(left) > 0 {
			res = append(res, left...)
		}
		if len(right) > 0 {
			res = append(res, right...)
		}
		return res
	}(mergeSort(nums[:mid]), mergeSort(nums[mid:]))
}

func main() {
	nums := getRandArray(20)
	fmt.Println(nums)
	fmt.Println(mergeSort(nums))

}
