package main

import "math"

func main() {

}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// 222. 完全二叉树的节点个数
// https://leetcode.cn/problems/count-complete-tree-nodes/
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left, right := root, root
	// 左边的高度
	lh := 0
	for left != nil {
		lh++
		left = left.Left
	}

	// 右边的高度
	rh := 0
	for right != nil {
		rh++
		right = right.Right
	}

	// 满二叉树
	if lh == rh {
		return int(math.Pow(2,float64(lh))) - 1
	}

	// 不是满二叉树，只能迭代求值
	return countNodes(root.Left) + countNodes(root.Right) + 1
}

