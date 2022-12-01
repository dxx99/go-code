package main
// 24. 两两交换链表中的节点
// 给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。
//
// 
//
//示例 1：
//
//
//输入：head = [1,2,3,4]
//输出：[2,1,4,3]
//示例 2：
//
//输入：head = []
//输出：[]
//示例 3：
//
//输入：head = [1]
//输出：[1]
// 
//
//提示：
//
//链表中节点的数目在范围 [0, 100] 内
//0 <= Node.val <= 100
//
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/swap-nodes-in-pairs
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {

}

type ListNode struct {
	Val int
	Next *ListNode
}


func swapPairs(head *ListNode) *ListNode {
	cur := head
	for cur != nil && cur.Next != nil {
		next := cur.Next
		cur.Next = next.Next
		next.Next = cur


		cur = next.Next
	}

	return head
}

func swapPairsV1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 栈空间存储变量
	newHead := head.Next
	head.Next = swapPairsV1(newHead.Next)

	// 后序遍历的位置
	newHead.Next = head
	return newHead
}