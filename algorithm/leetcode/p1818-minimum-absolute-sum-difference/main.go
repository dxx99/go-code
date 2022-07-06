package main

import (
	"fmt"
	"sort"
)

func main() {
	num1 := []int{53,48,14,71,31,55,6,80,28,19,15,40,7,21,69,15,5,42,86,15,11,54,44,62,9,100,2,26,81,87,87,18,45,29,46,100,20,87,49,86,14,74,74,52,52,60,8,25,21,96,7,90,91,42,32,34,55,20,66,36,64,67,44,51,4,46,25,57,84,23,10,84,99,33,51,28,59,88,50,41,59,69,59,65,78,50,78,50,39,91,44,78,90,83,55,5,74,96,77,46}
	num2 := []int{39,49,64,34,80,26,44,3,92,46,27,88,73,55,66,10,4,72,19,37,40,49,40,58,82,32,36,91,62,21,68,65,66,55,44,24,78,56,12,79,38,53,36,90,40,73,92,14,73,89,28,53,52,46,84,47,51,31,53,22,24,14,83,75,97,87,66,42,45,98,29,82,41,36,57,95,100,2,71,34,43,50,66,52,6,43,94,71,93,61,28,84,7,79,23,48,39,27,48,79}
	fmt.Println(minAbsoluteSumDiff(num1, num2))
	//fmt.Println(minAbsoluteSumDiff([]int{1,10,4,4,2,7}, []int{9,3,5,1,7,4}))
	//fmt.Println(minAbsoluteSumDiff([]int{1,7,5}, []int{2,3,5}))
	//fmt.Println(minAbsoluteSumDiff([]int{2,4,6,8,10}, []int{2,4,6,8,10}))
}

func minAbsoluteSumDiff(nums1 []int, nums2 []int) int {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	// 定义排序的num1
	sNum1 := make([]int, len(nums1))
	copy(sNum1, nums1)
	sort.Ints(sNum1)
	// 遍历num2, 做二分查询
	sum, maxN := 0, 0
	for i := 0; i < len(nums2); i++ {
		diff := abs(nums1[i] - nums2[i])
		sum += diff

		x := sort.SearchInts(sNum1, nums2[i])

		//todo 必须处理左右两边的元素
		if x < len(nums1) {
			maxN = max(maxN, diff-(sNum1[x]-nums2[i]))
		}
		if x > 0 {
			maxN = max(maxN, diff-(nums2[i] - sNum1[x-1]))
		}
	}

	return (sum - maxN) % (1e9+7)
}
