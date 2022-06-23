package main


func main() {

}


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func findBottomLeftValue(root *TreeNode) int {
	left, max, deep := 0, -1, 0
	var helper func(node *TreeNode)
	helper = func(node *TreeNode) {
		if node == nil {
			return
		}
		deep++
		if node.Left == nil && node.Right == nil && deep > max {
			max = deep
			left = node.Val
		}
		helper(node.Left)
		helper(node.Right)
		deep--
	}

	helper(root)

	return left
}