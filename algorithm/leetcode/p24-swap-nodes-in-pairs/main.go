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

func swapPairs(head *ListNode) *ListNode {
	dummyNode := &ListNode{0, head}

	tmp := dummyNode
	for tmp.Next != nil && tmp.Next.Next != nil {
		cur := tmp.Next
		next := tmp.Next.Next

		// 节点交换
		tmp.Next = next
		cur.Next = next.Next
		next.Next = cur

		// 遍历指针向前
		tmp = cur
	}

	return dummyNode.Next
}
