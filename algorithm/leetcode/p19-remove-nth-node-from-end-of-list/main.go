package main

//p19 删除链表的倒数第N个结点
// 给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
//
// 
//
//示例 1：
//
//
//输入：head = [1,2,3,4,5], n = 2
//输出：[1,2,3,5]
//示例 2：
//
//输入：head = [1], n = 1
//输出：[]
//示例 3：
//
//输入：head = [1,2], n = 1
//输出：[1]
// 
//
//提示：
//
//链表中结点的数目为 sz
//1 <= sz <= 30
//0 <= Node.val <= 100
//1 <= n <= sz
// 
//
//进阶：你能尝试使用一趟扫描实现吗？
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {

}

type ListNode struct {
	Val int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	l := 0
	tNode, ansNode := head, head
	for tNode != nil {
		l++
		tNode = tNode.Next
	}
	dIndex := l-n+1   // 删除第几个元素的索引
	if dIndex <= 0 {
		return ansNode
	}

	// 删除第一个元素
	if dIndex == 1 {
		ansNode = head.Next
		return ansNode
	}

	// 删除除第一个元素的任意一个
	dIndex--
	for head != nil && dIndex != 1 {
		dIndex--
		head = head.Next
	}
	head.Next = head.Next.Next

	return ansNode
}

// 代码优化
func removeNthFromEndV2(head *ListNode, n int) *ListNode {
	// 获取链表长度
	l := 0
	lNode := head
	for lNode != nil {
		l++
		lNode = lNode.Next
	}

	// 添加一个头节点，方便操作
	ansNode := &ListNode{
		Val:  0,
		Next: head,
	}
	cut := ansNode

	// 删除节点K, 也就是(l-n+1)
	for i := 0; i < l - n; i++ {
		cut = cut.Next
	}
	cut.Next = cut.Next.Next

	return ansNode.Next

}

// 利用栈求解
func removeNthFromEndV3(head *ListNode, n int) *ListNode {
	stack := make([]*ListNode, 0)
	ansNode := &ListNode{0, head}

	// 将数据全部压到栈内
	for node := ansNode; node != nil; node = node.Next {
		stack = append(stack, node)
	}

	// 删掉倒数第n个元素, 也就是 k = len(stack)-n 个元素
	// 去他前面的一个元素，然后把他后面的元素删掉
	pre := stack[len(stack)-1-n]
	pre.Next = pre.Next.Next
	return ansNode.Next
}

// 快慢指针
// 快指针先走n步，然后再慢指针一起走
func removeNthFromEndV4(head *ListNode, n int) *ListNode  {
	// 先包装一个头指针
	ansNode := &ListNode{0, head}

	// 快指针先走n步
	fast, slow := head, ansNode
	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	// 快慢指针一起走，知道快指针走完，这有有一个前开后闭的原则，也就是
	for ; fast != nil; fast = fast.Next {
		slow = slow.Next
	}

	// 删除慢指针的后一个元素
	slow.Next = slow.Next.Next
	return ansNode.Next
}
