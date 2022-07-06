package main

import (
	"fmt"
)

func main() {
	fmt.Println(maximumsSplicedArray([]int{28,34,38,14,30,31,23,7,28,3}, []int{42,35,7,6,24,30,14,21,20,34}))
	//fmt.Println(maximumsSplicedArray([]int{60,60,60}, []int{10,90,10}))
	//fmt.Println(maximumsSplicedArray([]int{20,40,20,70,30}, []int{50,20,50,40,20}))
	//fmt.Println(maximumsSplicedArray([]int{7,11,13}, []int{1,1,1}))
	//fmt.Println(countHousePlacements(1000))	// 500478595
	//fmt.Println(checkXMatrix([][]int{{2,0,0,1},{0,3,1,0},{0,5,2,0},{4,0,0,2}}))
}

//1
func checkXMatrix(grid [][]int) bool {
	n := len(grid)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if i == j || i+j == n-1 {
				if grid[i][j] == 0 {
					return false
				}
			}else {
				if grid[i][j] != 0 {
					return false
				}
			}
		}
	}

	return true
}


//2
func countHousePlacements(n int) int {
	if n == 1 {
		return 4
	}
	const mod = 1e9 + 7
	dp := make([]int, n+1)
	dp[1], dp[2] = 2, 3
	for i := 3; i <= n; i++ {
		dp[i] = (dp[i-1]+dp[i-2]) % mod
	}
	return (dp[n] * dp[n]) % mod
}


//3
func maximumsSplicedArray(nums1 []int, nums2 []int) int {
	t1, t2 := 0, 0
	for i := 0; i < len(nums1); i++ {
		t1 += nums1[i]
		t2 += nums2[i]
	}
	getSub := func(n1 []int, n2 []int) int {
		max, cur := 0, 0
		for i := 0; i < len(n1); i++ {
			cur += n2[i] - n1[i]
			if cur < 0 {
				cur = 0
			}
			if cur > max {
				max = cur
			}
		}
		return max
	}

	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	// 那边大就以那边为基础
	m1 := t1+getSub(nums1, nums2)
	m2 := t2 + getSub(nums2, nums1)
	return max(m1, m2)
}


//4
func minimumScore(nums []int, edges [][]int) int {
	return 0
}
