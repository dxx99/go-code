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
	Val  int
	Next *ListNode
}

func numComponents(head *ListNode, nums []int) int {
	m := make(map[int]bool)
	for _, num := range nums {
		m[num] = true
	}

	ans := 0
	flag := false
	for head != nil {
		if ok, _ := m[head.Val]; ok {
			flag = true
		} else {
			if flag {
				ans++
				flag = false
			}
		}
		head = head.Next
	}
	if flag {
		ans++
	}
	return ans
}
