package main

import "fmt"

func main() {
	fmt.Println(totalSteps([]int{10,1,2,3,4,5,6,1,2,3}))
	//fmt.Println(totalSteps([]int{5,3,4,4,7,3,6,11,8,5,11}))
	//fmt.Println(totalSteps([]int{4,5,7,7,13}))
}

func totalSteps(nums []int) int {
	ans := 0
	left, mid, right := 0, 0, 1
	for i := 1; i < len(nums); i++ {
		if nums[left] > nums[i] {
			mid = i
		}
		if nums[i] >= nums[i-1] {
			right = i
		}

		// 开始计算左右区间的值
		if nums[i] < nums[i-1] {
			v2 := right - mid
			fmt.Println( nums[right], nums[left])

			if nums[right] < nums[left] {
				v2 = right-mid+1
			}
			ans = max(ans, max(v2, mid-left))

			left, mid, right = i, i, i+1
		}
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
