package main

import (
	"fmt"
)

// p116 填充每个节点的下一个右侧节点的指针
// 给定一个 完美二叉树 ，其所有叶子节点都在同一层，每个父节点都有两个子节点。二叉树定义如下：
//
//struct Node {
//  int val;
//  Node *left;
//  Node *right;
//  Node *next;
//}
//填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL。
//
//初始状态下，所有 next 指针都被设置为 NULL。
//
// 
//
//示例 1：
//
//
//
//输入：root = [1,2,3,4,5,6,7]
//输出：[1,#,2,3,#,4,5,6,7,#]
//解释：给定二叉树如图 A 所示，你的函数应该填充它的每个 next 指针，以指向其下一个右侧节点，如图 B 所示。序列化的输出按层序遍历排列，同一层节点由 next 指针连接，'#' 标志着每一层的结束。
//示例 2:
//
//输入：root = []
//输出：[]
// 
//
//提示：
//
//树中节点的数量在 [0, 212 - 1] 范围内
//-1000 <= node.val <= 1000
// 
//
//进阶：
//
//你只能使用常量级额外空间。
//使用递归解题也符合要求，本题中递归程序占用的栈空间不算做额外的空间复杂度。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	node1 := &Node{
		Val:   1,
		Left:  nil,
		Right: nil,
		Next:  &Node{
			Val:   2,
			Left:  nil,
			Right: nil,
			Next:  nil,
		},
	}

	node2 := node1

	node2.Next = &Node{
		Val:   3,
		Left:  nil,
		Right: nil,
		Next:  nil,
	}

	fmt.Println(node2.Next.Val, node1.Next.Val)

	fmt.Printf("node1.Next = %p, node2.Next = %p", node1.Next, node2.Next)
}

type Node struct {
	Val int
	Left *Node
	Right *Node
	Next *Node
}

// todo 广度优先遍历二叉树，求值
func connect(root *Node) *Node {
	// 临界条件
	if root == nil {
		return root
	}
	queue := make([]*Node, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		tmp := queue
		queue = nil

		// 每层进行一次循环
		for i, node := range tmp {

			// 连接每一层的节点
			if i+1 < len(tmp) {
				node.Next = tmp[i+1]
			}

			// 往队列中插入元素
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return root
}
