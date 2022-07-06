package main

import (
	"fmt"
	"sort"
)

func main() {
	// [-2,-1,1,2,3,4,5]
	//7
	//3
	fmt.Println(findClosestElements([]int{-2,-1,1,2,3,4,5}, 7, 3))
	fmt.Println(findClosestElements([]int{0,0,1,2,3,3,4,7,7,8}, 3, 5))
	fmt.Println(findClosestElements([]int{0,0,0,1,3,5,6,7,8,8}, 2, 2))
	fmt.Println(findClosestElements([]int{1,2,5,5,6,6,7,7,8,9}, 7, 7))
	fmt.Println(findClosestElements([]int{0,0,1,2,3,3,4,7,7,8}, 3, 5))
	fmt.Println(findClosestElements([]int{1,2,3,4,5}, 4, 3))
	fmt.Println(findClosestElements([]int{1,2,3,4,5}, 4, 5))
	// output:
	// [-2 -1 1 2 3 4 5]
	// [3,3,4]
	// [1,3]
	// [5 5 6 6 7 7 8]
	// [3,3,4]
	// [1 2 3 4]
	// [2 3 4 5]
}

// 658. 找到 K 个最接近的元素
// https://leetcode.cn/problems/find-k-closest-elements/
func findClosestElements(arr []int, k int, x int) []int {
	p := sort.SearchInts(arr, x)
	left, right := p-k, p+k
	if left < 0 {
		left = 0
	}
	if right > len(arr)-1 {
		right = len(arr)-1
	}

	for right - left + 1 > k {
		if arr[right]-x >= x - arr[left] {
			right--
		}else {
			left++
		}
	}

	return arr[left:right+1]
}
