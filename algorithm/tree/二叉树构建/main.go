package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(constructMaximumBinaryTree([]int{3,2,1,6,0,5}))
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// 106. 从中序与后序遍历序列构造二叉树
// https://leetcode.cn/problems/construct-binary-tree-from-inorder-and-postorder-traversal/
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}

	v := postorder[len(postorder)-1]
	node := &TreeNode{Val: v}
	mIndex := 0
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == v {
			mIndex = i
			break
		}
	}

	node.Left = buildTree(inorder[:mIndex], postorder[:mIndex])
	node.Right = buildTree(inorder[mIndex+1:], postorder[mIndex:len(postorder)-1])

	return node
}

// 105. 从前序与中序遍历序列构造二叉树
// https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
func buildTreePre(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	v := preorder[0]
	node := &TreeNode{Val: v}
	mIndex := 0
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == v {
			mIndex = i
			break
		}
	}

	node.Left = buildTreePre(preorder[1:mIndex+1], inorder[:mIndex])
	node.Right = buildTreePre(preorder[mIndex+1:], inorder[mIndex+1:])

	return node
}

// 654. 最大二叉树
// https://leetcode.cn/problems/maximum-binary-tree/
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	max := math.MinInt
	mIndex := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
			mIndex = i
		}
	}
	root := &TreeNode{Val: max}

	root.Left = constructMaximumBinaryTree(nums[:mIndex])
	root.Right = constructMaximumBinaryTree(nums[mIndex+1:])

	return root
}

// 617.合并二叉树
// https://leetcode.cn/problems/merge-two-binary-trees/
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	}
	v := 0
	var left1, left2, right1, right2 *TreeNode
	if root1 != nil {
		v += root1.Val
		left1, right1 = root1.Left, root1.Right

	}
	if root2 != nil {
		v += root2.Val
		left2, right2 = root2.Left, root2.Right

	}
	root := &TreeNode{Val: v}

	root.Left = mergeTrees(left1, left2)
	root.Right = mergeTrees(right1, right2)

	return root
}