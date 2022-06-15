package main

func main() {

}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


// 236. 二叉树的最近公共祖先
// https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree/
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 终止条件
	if root == p || root == q || root == nil {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}


	if left == nil && right != nil {
		return right
	}
	if right == nil && left != nil {
		return left
	}
	return nil
}

// 235. 二叉搜索树的最近公共祖先
// https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-search-tree/
func lowestCommonAncestorV2(root, p, q *TreeNode) *TreeNode {
	var helper func(cur, p, q *TreeNode) *TreeNode

	helper = func(cur, p, q *TreeNode) *TreeNode {
		if cur == nil {
			return nil
		}

		// 如果当前节点大于q、p两值
		if cur.Val > p.Val && cur.Val > q.Val {
			left := helper(cur.Left, p, q)
			if left != nil {
				return left
			}
		}

		// 如果当前节点小于q、p两值
		if cur.Val < p.Val && cur.Val < q.Val {
			right := helper(cur.Right, p, q)
			if right != nil {
				return right
			}
		}
		return cur
	}

	return helper(root, p, q)
}