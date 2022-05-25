package main

//定义一个函数，输入一个链表的头节点，反转该链表并输出反转后链表的头节点。
//
// 
//
//示例:
//
//输入: 1->2->3->4->5->NULL
//输出: 5->4->3->2->1->NULL
// 
//
//限制：
//
//0 <= 节点个数 <= 5000
//
//
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/fan-zhuan-lian-biao-lcof
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

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var prev *ListNode
	cur := head
	for cur != nil {
		next := cur.Next

		cur.Next = prev
		prev = cur

		cur = next
	}

	return prev
}

func reverseListV2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	root := &ListNode{
		Val:  0,
		Next: nil,
	}

	for head != nil {
		tmp := root.Next
		root.Next = &ListNode{
			Val:  head.Val,
			Next: nil,
		}
		root.Next.Next = tmp

		head = head.Next
	}

	return root.Next
}
