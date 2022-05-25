package main

//输入一个链表，输出该链表中倒数第k个节点。为了符合大多数人的习惯，本题从1开始计数，即链表的尾节点是倒数第1个节点。
//
//例如，一个链表有 6 个节点，从头节点开始，它们的值依次是 1、2、3、4、5、6。这个链表的倒数第 3 个节点是值为 4 的节点。
//
// 
//
//示例：
//
//给定一个链表: 1->2->3->4->5, 和 k = 2.
//
//返回链表 4->5.
//
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/lian-biao-zhong-dao-shu-di-kge-jie-dian-lcof
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

func getKthFromEnd(head *ListNode, k int) *ListNode {
	if k <= 0 {
		return nil
	}
	ans := make([]*ListNode, 0)

	for head != nil {
		ans = append(ans, head)
		head = head.Next
	}

	if k > len(ans) {
		return ans[0]
	}

	return ans[len(ans)-k]
}

func getKthFromEndV2(head *ListNode, k int) *ListNode {
	if k <= 0 {
		return nil
	}

	// 快慢指针求解
	fast, slow := head, head
	for fast != nil && k > 0 {
		fast = fast.Next
		k--
	}

	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	return slow
}


