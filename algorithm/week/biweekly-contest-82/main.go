package main

import (
	"fmt"
	"math"
	"sort"
)

// https://leetcode.cn/contest/biweekly-contest-82
func main() {
	//fmt.Println(minSumSquareDiff([]int{1}, []int{5},1,1))
	fmt.Println(minSumSquareDiff([]int{1,4,10,12}, []int{5,8,6,9}, 1,1))
	//fmt.Println(minSumSquareDiff([]int{10,10,10,11,5}, []int{1,0,6,6,1}, 11,27))
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

//1
func evaluateTree(root *TreeNode) bool {
	if root.Left == nil && root.Right == nil {
		if root.Val == 0 {
			return false
		}else {
			return true
		}
	}
	left := evaluateTree(root.Left)
	right := evaluateTree(root.Right)

	if root.Val == 2 {
		return left || right
	}
	return left && right
}


//2


//3
func minSumSquareDiff(nums1 []int, nums2 []int, k1 int, k2 int) int64 {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}


	arr := make([]int, len(nums1))
	for i := 0; i < len(arr); i++ {
		arr[i] = abs(nums2[i] - nums1[i])
	}
	// 排序
	sort.Ints(arr)
	k := k1+k2
	fmt.Println(arr)

	// 特殊情况
	tmp := 0
	for i := 0; i < len(arr); i++ {
		tmp += arr[i]
	}
	if k >= tmp {
		return 0
	}

	f := func(mid int) bool {
		sum := 0
		v := arr[mid]
		for i := mid+1; i < len(arr); i++ {
			sum += arr[i] - v
		}
		return sum > k
	}

	// 每个元素减x
	left, right := 0, len(arr)-1
	for left < right {
		mid := int(uint(left+right)>>1)
		if f(mid) {
			left = mid+1
		}else {
			right = mid
		}
	}
	fmt.Println(left)


	var sum int64
	for i := 0; i < len(arr); i++ {
		if i >= left {
			sum += int64(math.Pow(float64(arr[left]), 2))
		}else {
			sum += int64(math.Pow(float64(arr[i]), 2))
		}
	}

	return sum
}


//4

