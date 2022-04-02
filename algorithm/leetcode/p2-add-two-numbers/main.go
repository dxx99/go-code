package main

import (
	"fmt"
	"math"
)

// p2 两数相加
// 给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
//
//请你将两个数相加，并以相同形式返回一个表示和的链表。
//
//你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
//
// 
//
//示例 1：
//
//
//输入：l1 = [2,4,3], l2 = [5,6,4]
//输出：[7,0,8]
//解释：342 + 465 = 807.
//示例 2：
//
//输入：l1 = [0], l2 = [0]
//输出：[0]
//示例 3：
//
//输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
//输出：[8,9,9,9,0,0,0,1]
// 
//
//提示：
//
//每个链表中的节点数在范围 [1, 100] 内
//0 <= Node.val <= 9
//题目数据保证列表表示的数字不含前导零
//
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/add-two-numbers
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	num := 123

	tArr := make([]int, 0)
	for num >= 10 {
		tArr = append(tArr, num%10)
		num = num / 10
	}
	tArr = append(tArr, num)

	fmt.Println(int( math.Pow10(0)))
	fmt.Println(int( math.Pow10(1)))
}

type ListNode struct {
	Val int
	Next *ListNode
}


// 思路比较复杂
// 需要先转化成数字，再相加，再转成链表
// #BUG 会出现溢出问题，由于数据太长
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var num1, num2 int

	pre := 0
	for l1 != nil || l2 != nil {
		if l1 != nil {
			num1 += l1.Val * int( math.Pow10(pre))
			l1 = l1.Next
		}
		if l2 != nil {
			num2 += l2.Val * int( math.Pow10(pre))
			l2 = l2.Next
		}
		pre++
	}

	tSum := num1 + num2
	fmt.Printf("%d + %d = %d\n", num1, num2, tSum)

	//解析成个位元素
	tArr := make([]int, 0)
	for tSum >= 10 {
		tArr = append(tArr, tSum%10)
		tSum = tSum / 10
	}
	tArr = append(tArr, tSum)

	// 遍历存储到节点中
	head := new(ListNode)
	if len(tArr) == 0 {
		return head
	}

	tNode := head
	tNode.Val = tArr[0]
	for i := 1; i < len(tArr); i++ {
		node := &ListNode{
			Val:  tArr[i],
			Next: nil,
		}
		tNode.Next = node
		tNode = tNode.Next

	}

	return head
}

