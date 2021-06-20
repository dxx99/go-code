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

// 快速排序：通过一趟排序将待排序记录分隔成独立的两部分，其中一部分记录关键字均比另一部分的关键字小，然后分别对这两部分记录进行排序，以达到整个序有序
// 算法关键字： 从数组中挑选基准值，
// 算法描述：
// 1. 从数列中挑出一个元素，称为 “基准”（pivot）；
// 2. 重新排序数列，所有元素比基准值小的摆放在基准前面，所有元素比基准值大的摆在基准的后面（相同的数可以到任一边）。在这个分区退出之后，该基准就处于数列的中间位置。这个称为分区（partition）操作；
// 3. 递归地（recursive）把小于基准值元素的子数列和大于基准值元素的子数列排序。
func quickSort(nums []int , left, right int) []int {
	if left < right {
		pIndex := partition(nums, left, right)
		quickSort(nums, left, pIndex - 1)
		quickSort(nums, pIndex + 1, right)
	}
	return nums
}

// 返回分区临界点
func partition(nums []int, left, right int) int {
	pivot := left
	pIndex := pivot +1
	for i := pIndex; i <= right; i++ {
		// 基准元素与要比较的元素进行比较，要比较的元素小于基准元素，则进行元素交换, 基准元素的索引加一
		if nums[i] < nums[pivot] {
			// 交换操作
			nums[i], nums[pIndex] = nums[pIndex], nums[i]
			pIndex++
		}
	}

	// 基准元素位置交换
	nums[pivot], nums[pIndex-1] = nums[pIndex-1], nums[pivot]
	return pIndex-1
}



// 快速排序：
func main() {
	nums := getRandArray(10)
	fmt.Println(nums)
	fmt.Println(quickSort(nums, 0, len(nums)-1))

}
