package main

import (
	"fmt"
	"math/rand"
)

// 网易
// https://leetcode.cn/study-plan/zhitongche/?progress=49tbnn3
func main() {
	//fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))

	fmt.Println(twoSum([]int{3, 2, 4}, 6))
}

// ------------------------第一题------------------------------//
// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/
func maxProfit(prices []int) int {
	cur, max := 0, 0
	for i := 1; i < len(prices); i++ {
		if prices[i]-prices[i-1]+cur < 0 {
			cur = 0
		} else {
			cur += prices[i] - prices[i-1]
		}
		if cur > max {
			max = cur
		}
	}
	return max
}

// --------------------------------------第二题---------------------------------------------//
// https://leetcode.cn/problems/linked-list-random-node/

type ListNode struct {
	Val  int
	Next *ListNode
}

type Solution struct {
	arr []int
	k   int
}

func Constructor(head *ListNode) Solution {
	arr := make([]int, 0)
	k := 0
	for head != nil {
		k++
		arr = append(arr, head.Val)
		head = head.Next
	}
	return Solution{
		arr: arr,
		k:   k,
	}
}

func (s *Solution) GetRandom() int {
	k := rand.Intn(s.k)
	return s.arr[k]
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */

// -----------------------------------第三题-----------------------------------------
// https://leetcode.cn/problems/two-sum/?envType=study-plan&id=wangyi&plan=zhitongche&plan_progress=49tbnn3
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for x, item := range nums {
		if y, ok := m[target-item]; ok {
			return []int{y, x}
		}
		m[item] = x
	}

	return []int{}
}
