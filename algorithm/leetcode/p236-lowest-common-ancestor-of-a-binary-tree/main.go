package main

func main() {

}

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == p || root == q || root == nil {
		return root
	}
	
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	// 后续遍历的位置
	if left != nil && right != nil {
		return root
	}
	
	// 返回左值
	if left != nil {
		return left
	}
	// 返回右值
	if right != nil {
		return right
	}
	return nil
}
