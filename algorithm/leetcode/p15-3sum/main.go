package main

import (
	"fmt"
	"sort"
)

// p15 三数之和
func main() {
	fmt.Println(threeSum([]int{-1,0,1,2,-1,-4}))
	fmt.Println(threeSum([]int{}))
	fmt.Println(threeSum([]int{0}))
}

func threeSum(nums []int) [][]int {
	l := len(nums)
	res := make([][]int,0)

	if l < 3 {
		return res
	}
	// 先排序
	sort.Ints(nums)

	for i := 0; i < l-1; i++ {
		// 重复元素不用两次处理
		if i >0 && nums[i] == nums[i-1] {
			continue
		}
		
		left, right := i+1, l-1
		for left < right {
			if nums[i] + nums[left] + nums[right] == 0 {
				res = append(res, []int{nums[i], nums[left], nums[right]})

				// 解决重复元素问题
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if nums[i] + nums[left] + nums[right] > 0 {
				right--
			} else {
				left++
			}
		}
	}
	return res
}
