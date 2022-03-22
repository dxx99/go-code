package main

import "math"

//给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。
//
//有效 二叉搜索树定义如下：
//
//节点的左子树只包含 小于 当前节点的数。
//节点的右子树只包含 大于 当前节点的数。
//所有左子树和右子树自身必须也是二叉搜索树。
// 
// 示例一： https://assets.leetcode.com/uploads/2020/12/01/tree1.jpg
//输入：root = [2,1,3]
//输出：true
//
// 示例二： https://assets.leetcode.com/uploads/2020/12/01/tree2.jpg
// 输入：root = [5,1,4,null,null,3,6]
//输出：false
//解释：根节点的值是 5 ，但是右子节点的值是 4 。
//
// 提示：
//树中节点数目范围在[1, 104] 内
//-231 <= Node.val <= 231 - 1
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/validate-binary-search-tree
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {

}

// TreeNode Definition for a binary tree node.
// * type TreeNode struct {
// *     Val int
// *     Left *TreeNode
// *
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// 1. 递归，要有终止条件，就是nil节点
// 2. 对每一个节点进行判断，设置上下界
func isValidBST(root *TreeNode)  bool {
	return helper(root, nil, nil)
}

func helper(root *TreeNode, lower *TreeNode, upper *TreeNode) bool {
	if root == nil {
		return true
	}

	// 小于下界
	if lower != nil && root.Val <= lower.Val {
		return false
	}

	// 大于上界
	if upper != nil && root.Val >= upper.Val {
		return false
	}

	// 左边的节点只有可能越上界，右边的节点只有可能越下界
	return helper(root.Left, lower, root) && helper(root.Right, root, upper)
}


// 中序遍历解法
func isValidBSTV2(root *TreeNode) bool {
	stack := make([]*TreeNode, 0)
	inOrderNum := math.MinInt64

	for len(stack) > 0 || root != nil {

		// 先把左节点压入栈
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		// 出栈
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		//比接
		if root.Val <= inOrderNum {
			return false
		}

		// 替换
		inOrderNum = root.Val
		root = root.Right
	}
	return true
}



