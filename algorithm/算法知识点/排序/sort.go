package main

import (
	"fmt"
	"sort"
)

func main() {
	sort.Ints([]int{})
	fmt.Println(quickSort([]int{1,5,2,11,5,3,89,5,3,21,77}, 0, 10))
}

// 归并排序：核心思想就是"二叉树的后序遍历"
func mergeSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	merge := func(nums1, nums2 []int) []int {
		n1, n2 := len(nums1), len(nums2)
		res := make([]int, n1+n2)

		k := 0
		k1, k2 := 0, 0
		for k1 < n1 && k2 < n2 {
			if nums1[k1] > nums2[k2] {
				res[k] = nums2[k2]
				k2++
			}else {
				res[k] = nums1[k1]
				k1++
			}
			k++
		}
		for i := k1; i < n1; i++ {
			res[k] = nums1[i]
			k++
		}
		for j := k2; j < n2; j++ {
			res[k] = nums2[j]
			k++
		}
		return res
	}

	mid := len(nums)>>1
	left := mergeSort(nums[:mid])
	right := mergeSort(nums[mid:])

	//后序操作，也就是合并操作
	return merge(left, right)
}

func mergeV1(left, right []int) []int {
	res := make([]int, 0)
	for len(left) > 0 && len(right) > 0 {
		if left[0] > right[0] {
			res = append(res, right[0])
			right = right[1:]
		}else {
			res = append(res, left[0])
			left = left[1:]
		}
	}
	//左右元素合并
	if len(left) > 0 {
		res = append(res, left...)
	}
	if len(right) > 0 {
		res = append(res, right...)
	}
	return res
}



// 快速排序：核心思想就是"二叉树的前序遍历"
func quickSort(nums []int, left, right int) []int {
	// 选取一个基准点，将数据分成两边
	partArr := func(left, right int) int {
		pivot := left	// 这个就是临界点
		px := pivot+1
		// 注意区间，这里都是左边右闭区间，所以要等于
		for i := px; i <= right; i++ {
			if nums[i] < nums[pivot] {
				nums[i], nums[px] = nums[px], nums[i]
				px++
			}
		}
		nums[pivot], nums[px-1] = nums[px-1], nums[pivot]
		return px-1
	}

	if left < right {
		// 前序位置
		pIndex := partArr(left, right)

		// 注意区间位置，因为pIndex这个位置已经排好了
		quickSort(nums, left, pIndex-1)
		quickSort(nums, pIndex+1, right)
	}
	return nums
}
