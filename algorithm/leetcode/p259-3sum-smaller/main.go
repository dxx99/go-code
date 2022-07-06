package main

import (
	"fmt"
	"sort"
)

func main() {
	//fmt.Println(threeSumSmaller([]int{-2,0,1,3}, 2))		//2
	//fmt.Println(threeSumSmaller([]int{-2,0,1,4}, 2))	 	//1
	fmt.Println(threeSumSmaller([]int{3,1,0,-2}, 4))	//3
	fmt.Println(threeSumSmallerV2([]int{3,1,0,-2}, 4))	//3
	//fmt.Println(threeSumSmaller([]int{}, 0))
	//fmt.Println(threeSumSmaller([]int{0}, 0))
}

// 259. 较小的三数之和
// 三数之后小于target组合个数，
// 链接：https://leetcode.cn/problems/3sum-smaller
func threeSumSmaller(nums []int, target int) int {
	ans:= 0
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		k := len(nums)-1
		for j := i+1; j < len(nums)-1; j++ {
			for k > j && nums[i]+nums[j]+nums[k] >= target {
				k--
			}
			//
			fmt.Println("k-->", k, "j-->", j)
			if k - j > 0 {
				ans += k - j
			}
		}
	}
	return ans
}

// 三数之和小于某个值，可以转换成两数之和小于
// 算法的性能为O(n^2) 高于上面的性能
func threeSumSmallerV2(nums []int, target int) int {
	ans:= 0
	sort.Ints(nums)

	// 两数之和小于，可以用双指针
	twoSumSmall := func(left int, t int) {
		right := len(nums)-1
		for left < right {
			if nums[left]+nums[right] < t {
				// 这里是临界条件, nums[left]和nums[left:right]之间的每个元素都小于目标值
				// 因为左右的变动就有多种情况
				ans += right - left
				left++	// 注意
			}else {
				right--
			}
		}
	}

	for i := 0; i < len(nums)-2; i++ {
		twoSumSmall(i+1, target-nums[i])
	}
	return ans
}


