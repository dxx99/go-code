package main

import "fmt"

func main() {
	fmt.Println(maxWidthRamp([]int{6,0,8,2,1,5}))
	fmt.Println(maxWidthRamp([]int{9,8,1,0,1,9,4,0,4,1}))
}




func maxWidthRamp(nums []int) int {
	maxRamp := 0
	for i := 0; i < len(nums)-1; i++ {
		start := nums[i]

		right := len(nums)-1
		for nums[right] < start && right > i {
			right--
		}

		if right - i > maxRamp {
			maxRamp = right - i
		}
	}
	return maxRamp
}
