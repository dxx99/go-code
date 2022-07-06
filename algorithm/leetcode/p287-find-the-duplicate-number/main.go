package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findDuplicate([]int{3,2,2,2,4}))	//4
	//fmt.Println(findDuplicate([]int{1,4,4,2,4}))	//4
	//fmt.Println(findDuplicate([]int{1,3,4,2,1}))
	//fmt.Println(findDuplicate([]int{1,3,4,2,2}))
	//fmt.Println(findDuplicate([]int{3,1,3,4,2}))
}

func findDuplicate(nums []int) int {
	sort.Ints(nums)
	left, right := 0, len(nums)-1

	for left < right {
		mid := int(uint(left+right)>>1)

		// cut: 小于中间值的元素个数
		cut := 0
		for i := 0; i < len(nums); i++ {
			if nums[i] <= mid {
				cut++
			}
		}

		if cut > mid {
			right = mid
		}else {
			left = mid+1
		}
	}
	return left
}


func findDuplicateV2(nums []int) int {
	slow, fast := 0, 0

	slow, fast = nums[slow], nums[nums[fast]]
	for  slow != fast {
		slow, fast = nums[slow], nums[nums[fast]]
	}

	slow = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}
