package main

import "fmt"

// p206 反转链表
// 给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
// 
//
//示例 1：
//
//
//输入：head = [1,2,3,4,5]
//输出：[5,4,3,2,1]
//示例 2：
//
//
//输入：head = [1,2]
//输出：[2,1]
//示例 3：
//
//输入：head = []
//输出：[]
// 
//
//提示：
//
//链表中节点的数目范围是 [0, 5000]
//-5000 <= Node.val <= 5000
// 
//
//进阶：链表可以选用迭代或递归方式完成反转。你能否用两种方法解决这道题？
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/reverse-linked-list
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	// init node list
	head := &ListNode{0, nil}
	tmp := head
	for i := 1; i <= 5; i++ {
		tmp.Next = &ListNode{i, nil}
		tmp = tmp.Next
	}

	resNodeList := reverseListV4(head.Next)

	for resNodeList != nil{
		fmt.Println(resNodeList)
		resNodeList = resNodeList.Next
	}

}

type ListNode struct {
	Val int
	Next *ListNode
}

// 反转链表
// use memory too much
func reverseList(head *ListNode) *ListNode {
	stack := make([]*ListNode, 0)

	for head != nil {
		stack = append(stack, head)
		head = head.Next
	}

	prev := &ListNode{0, nil}
	node := prev

	for i := len(stack)-1; i >= 0; i-- {
		node.Next = stack[i]
		node = node.Next
	}
	return prev.Next
}

// 迭代求解
func reverseListV2(head *ListNode) *ListNode {
	// 用来保存上一个值
	var prev *ListNode

	// 其实也就是链表的头插法
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev

		// 上一个值为当前值，当前值为下一个值
		prev, curr = curr, next
	}
	return prev
}

// 递归求解
func reverseListV3(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := reverseListV3(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

// 头插方求解
func reverseListV4(node *ListNode) *ListNode {
	head := &ListNode{0, nil}

	for node != nil {
		tmp := head.Next
		head.Next = &ListNode{node.Val, nil}
		head.Next.Next = tmp
		node = node.Next
	}

	return head.Next
}

