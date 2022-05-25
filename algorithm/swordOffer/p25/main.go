package main

//输入两个递增排序的链表，合并这两个链表并使新链表中的节点仍然是递增排序的。
//
//示例1：
//
//输入：1->2->4, 1->3->4
//输出：1->1->2->3->4->4
//限制：
//
//0 <= 链表长度 <= 1000
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/he-bing-liang-ge-pai-xu-de-lian-biao-lcof
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
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

// 定义头节点方便操作
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{
		Val:  0,
		Next: nil,
	}

	cur := head
	for l1 != nil || l2 != nil {
		// 为空情况特殊处理
		if l1 == nil {
			cur.Next = l2
			break
		}
		if l2 == nil {
			cur.Next = l1
			break
		}

		// 节点比较合并
		if l1.Val > l2.Val {
			cur.Next = l2
			l2 = l2.Next
		} else {
			cur.Next = l1
			l1 = l1.Next
		}
		cur = cur.Next
	}
	return head.Next
}


// 使用递归合并链表
func mergeTwoListsV2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val {
		l1.Next = mergeTwoListsV2(l1.Next, l2)
		return l1
	}else {
		l2.Next = mergeTwoListsV2(l1, l2.Next)
		return l2
	}
}

