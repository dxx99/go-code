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

func deleteNode(head *ListNode, val int) *ListNode {
	// 如果被删的元素在头节点，这需要替换节点
	if head.Val == val {
		return head.Next
	}

	// 被删的元素不在头节点
	pre := head
	for pre.Next != nil {
		if pre.Next.Val == val  {
			pre.Next = pre.Next.Next
			break
		}
		pre = pre.Next
	}
	return head
}
