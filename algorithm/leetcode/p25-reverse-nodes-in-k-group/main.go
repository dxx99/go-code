package main

import "fmt"

// 给你链表的头节点 head ，每 k 个节点一组进行翻转，请你返回修改后的链表。
//
//k 是一个正整数，它的值小于或等于链表的长度。如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
//
//你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。
//
// 
//
//示例 1：
//
//
//输入：head = [1,2,3,4,5], k = 2
//输出：[2,1,4,3,5]
//示例 2：
//
//
//
//输入：head = [1,2,3,4,5], k = 3
//输出：[3,2,1,4,5]
// 
//
//提示：
//链表中的节点数目为 n
//1 <= k <= n <= 5000
//0 <= Node.val <= 1000
// 
//
//进阶：你可以设计一个只用 O(1) 额外内存空间的算法解决此问题吗？
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/reverse-nodes-in-k-group
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {

	head := &ListNode{1, nil}
	head.Next = &ListNode{2, nil}

	newHead := reverseKGroup(head, 2)

	for newHead != nil {
		fmt.Printf("%d\t", newHead.Val)
		newHead = newHead.Next
	}


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

// 利用环形数组优化代码
func reverseKGroup(head *ListNode, k int) *ListNode {
	stack := make([]*ListNode, 0)
	dummy := &ListNode{0, nil}
	tmp := dummy
	c := k
	for head != nil {
		stack = append(stack, head)
		head = head.Next
		c--

		// 需要反转，出栈
		if c == 0 {
			for len(stack) != 0 {
				node := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				node.Next = nil
				tmp.Next = node
				tmp = tmp.Next
			}

			// 环形指针复原
			c = k
		}
	}

	// 处理栈中剩下的元素
	for _, node := range stack {
		node.Next = nil
		tmp.Next = node
		tmp = tmp.Next
	}

	return dummy.Next
}

//
