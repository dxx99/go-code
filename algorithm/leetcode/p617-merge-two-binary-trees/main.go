package main

import "fmt"

// p617 合并二叉树
//给你两棵二叉树： root1 和 root2 。
//
//想象一下，当你将其中一棵覆盖到另一棵之上时，两棵树上的一些节点将会重叠（而另一些不会）。你需要将这两棵树合并成一棵新二叉树。合并的规则是：如果两个节点重叠，那么将这两个节点的值相加作为合并后节点的新值；否则，不为 null 的节点将直接作为新二叉树的节点。
//
//返回合并后的二叉树。
//
//注意: 合并过程必须从两个树的根节点开始。
//
// 
//
//示例 1：
//
//
//输入：root1 = [1,3,2,5], root2 = [2,1,3,null,4,null,7]
//输出：[3,4,5,5,4,null,7]
//示例 2：
//
//输入：root1 = [1], root2 = [1,2]
//输出：[2,2]
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/merge-two-binary-trees
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {

	// 这种是不报错的
	queue := []int{1}
	fmt.Println(queue[0])
	fmt.Println(queue[1:])

}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	// 递归的终止条件，只要有其中一个节点为nil, 就直接取另外一个节点即可
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}

	root1.Val += root2.Val
	root1.Left = mergeTrees(root1.Left, root2.Left)
	root1.Right = mergeTrees(root1.Right, root2.Right)
	return root1
}

