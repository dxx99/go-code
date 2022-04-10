package main

import "fmt"

// p11 盛最多水的容器
func main() {
	fmt.Println(maxArea([]int{1,8,6,2,5,4,8,3,7}))
	fmt.Println(maxArea([]int{1,1}))
}

func maxArea(height []int) int {
	left, right := 0, len(height) -1
	max := 0
	for left < right {
		area := 0
		if height[right] > height[left] {
			area = height[left]*(right-left)
			left++
		}else {
			area = height[right]*(right-left)
			right--
		}
		if area > max {
			max = area
		}
	}

	return max
}
