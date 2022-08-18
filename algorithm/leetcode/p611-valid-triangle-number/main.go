package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(triangleNumberV2([]int{2}))
	fmt.Println(triangleNumberV2([]int{0,0,0}))
	fmt.Println(triangleNumberV2([]int{2,2,3,4}))
	fmt.Println(triangleNumberV2([]int{4,2,3,4}))
	//output:
	//0
	//0
	//3
	//4
	fmt.Println("v3.............")
	fmt.Println(triangleNumberV3([]int{2}))
	fmt.Println(triangleNumberV3([]int{0,0,0}))
	fmt.Println(triangleNumberV3([]int{2,2,3,4}))
	fmt.Println(triangleNumberV3([]int{4,2,3,4}))
}

// 611. 有效三角形的个数
// binarySearch O((n^2)*log(n))
// 链接地址: https://leetcode.cn/problems/valid-triangle-number/solution/
func triangleNumber(nums []int) int {
	ans := 0
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		for j := i+1; j < len(nums)-1; j++ {
			// 第一种，这里暴利顺序查找
			ans += sort.SearchInts(nums[j+1:], nums[i]+nums[j])		// 找到第一个等于该元素的索引位置
		}
	}
	return ans
}

// 双指针求解
// 性能优化： O(n^2)
func triangleNumberV2(nums []int) int {
	ans := 0
	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		k := i	// 记录每次k的位置，因为下次num[j]会增加，肯定会大于当前的临界值
		for j := i+1; j < len(nums)-1; j++ {
			for k+1 < len(nums) && nums[i] + nums[j] > nums[k+1] {
				k++
			}
			// 有可能出现负数，所以要控制
			if k - j > 0 {
				ans += k - j
			}
		}
	}

	return ans
}

// 全部转换成两数之和的问题
// 两数之和大于某一个值
// 性能更优，完全是O(n^2)
func triangleNumberV3(nums []int) int {
	ans := 0
	sort.Ints(nums)

	for i := len(nums)-1; i >= 2 ; i-- {
		// 节点都像中间靠拢
		left, right := 0, i-1
		for left < right {
			if nums[left]+nums[right] > nums[i] {
				ans += right - left
				right--
			}else {
				left++
			}
		}
	}

	return ans
}
