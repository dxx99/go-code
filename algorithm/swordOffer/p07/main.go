package main

import "fmt"

// 树的遍历
// 前序遍历：【根节点，左子树，右子树】
// 中序遍历：【左子树，根节点，右子树】
// 后续遍历：【左子树，右子树，根节点】
func main() {
	fmt.Println(buildTree([]int{3,9,20,15,7}, []int{9,3,15,20,7}))
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

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	// 根节点
	root := &TreeNode{
		Val:   preorder[0],
		Left:  nil,
		Right: nil,
	}

	// 找到根节点中序数组中的位置
	i := 0
	for i < len(inorder) {
		if inorder[i] == preorder[0] {
			break
		}
		i++
	}

	root.Left = buildTree(preorder[1:len(inorder[:i])+1], inorder[:i])
	root.Right = buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])

	return root
}
