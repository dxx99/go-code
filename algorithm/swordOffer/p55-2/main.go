package main

import "math"

// 输入一棵二叉树的根节点，判断该树是不是平衡二叉树。如果某二叉树中任意节点的左右子树的深度相差不超过1，那么它就是一棵平衡二叉树。
//
// 
//
//示例 1:
//
//给定二叉树 [3,9,20,null,null,15,7]
//
//    3
//   / \
//  9  20
//    /  \
//   15   7
//返回 true 。
//
//示例 2:
//
//给定二叉树 [1,2,2,3,3,null,null,4,4]
//
//       1
//      / \
//     2   2
//    / \
//   3   3
//  / \
// 4   4
//返回 false 。
//
// 
//
//限制：
//
//0 <= 树的结点个数 <= 10000
//注意：本题与主站 110 题相同：https://leetcode-cn.com/problems/balanced-binary-tree/
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/ping-heng-er-cha-shu-lcof
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
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


func isBalanced(root *TreeNode) bool {

	if root == nil {
		return true
	}
	return math.Abs(float64(height(root.Left) - height(root.Right))) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func height(node *TreeNode) int {
	if node == nil {
		return 0
	}

	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	return max(height(node.Left), height(node.Right)) +1
}