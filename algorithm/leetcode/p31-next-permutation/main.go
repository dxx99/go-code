package main

import (
	"fmt"
)

func main() {
	arr := []int{1,1}
	nextPermutation(arr)
	nextPermutationV2(arr)
	fmt.Println(arr)
}

// 先找到较小的数，再找较大的数
// 较小的数尽可能小，较大的数尽量靠右边
func nextPermutation(nums []int)  {
	// 得到较小的数
	i := len(nums) - 2
	for i >=0 && nums[i] >= nums[i+1] {
		i--
	}

	// 得到较大的数
	if i >= 0 {
		j := len(nums)-1
		for j >= 0 && nums[i] >= nums[j] {
			j--
		}
		fmt.Println(i, j)
		nums[i], nums[j] = nums[j], nums[i]
	}

	f := func(arr []int) {
		l, r := 0, len(arr)-1
		for l < r {
			arr[l], arr[r] = arr[r], arr[l]
			l++
			r--
		}
	}

	// 反转
	f(nums[i+1:])
}


func nextPermutationV2(nums []int) {
	n := len(nums)
	i := n - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	if i >= 0 {
		j := n - 1
		for j >= 0 && nums[i] >= nums[j] {
			j--
		}

		nums[i], nums[j] = nums[j], nums[i]
	}
	reverse(nums[i+1:])
}

func reverse(a []int) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}
