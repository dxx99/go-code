package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(purchasePlans([]int{3,1,2}, 5))
	//fmt.Println(purchasePlans([]int{2,5,3,5}, 6))
	//fmt.Println(purchasePlans([]int{2,2,1,9}, 10))
}

// 杨辉三角问题
func getRow(rowIndex int) []int {
	last := []int{1}
	ans := last
	for i := 0; i < rowIndex; i++ {
		ans = []int{last[0]}
		for j := 1; j < len(last); j++ {
			ans = append(ans, last[j]+last[j-1])
		}
		ans = append(ans, last[len(last)-1])
		last = ans
	}
	return ans
}

//
func purchasePlans(nums []int, target int) int {
	sort.Ints(nums)
	const  mod = 1e9+7
	ans := 0
	j := len(nums)-1
	for i := 0; i < len(nums); i++ {
		for ; j > i ; j++ {
			if nums[i]+nums[j] <= target {
				break
			}
		}
		if j > i {
			ans += j-i
		}
	}
	return ans%mod
}

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

type Node struct {
	Val int
	Children []*Node
}
var max = 0
func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	if len(root.Children) == 0 {
		return 1
	}
	for i := 0; i < len(root.Children); i++ {
		childNum := maxDepth(root.Children[i])+1
		if childNum > max {
			max = childNum
		}
	}
	return max
}