package main

func main() {
}


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	ans := make([]int, 0)

	for head != nil {
		ans = append(ans, head.Val)
		head = head.Next
	}

	// 反转数组
	left, right := 0, len(ans)-1
	for left < right {
		ans[left], ans[right] = ans[right], ans[left]
		left++
		right--
	}

	return ans
}