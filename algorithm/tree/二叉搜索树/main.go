package main

import (
	"math"
)

func main() {

}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// 700. 二叉搜索树中的搜索
// https://leetcode.cn/problems/search-in-a-binary-search-tree/
func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}else if root.Val < val {
		return searchBST(root.Right, val)
	}else {
		return searchBST(root.Left, val)
	}
}

// 98. 验证二叉搜索树
// https://leetcode.cn/problems/validate-binary-search-tree/
func isValidBST(root *TreeNode) bool {
	var helper func(node *TreeNode, lower *TreeNode, upper *TreeNode) bool
	helper = func(node *TreeNode, lower *TreeNode, upper *TreeNode) bool {
		if node == nil {
			return true
		}

		// 当前节点小于下界
		if lower != nil && node.Val <= lower.Val {
			return false
		}
		// 大于大于上界
		if upper != nil && node.Val >= upper.Val {
			return false
		}

		left := helper(node.Left, lower, node)		// 左边不要大于根节点
		right := helper(node.Right, node, upper)	// 右边不要小于根节点
		return left && right
	}

	return helper(root, nil, nil)
}


// 530. 二叉搜索树的最小绝对差
// https://leetcode.cn/problems/minimum-absolute-difference-in-bst/
// 看提示node.Val 是一个正数
func getMinimumDifference(root *TreeNode) int {
	ans := math.MaxInt

	last := math.MaxInt
	var helper func(node *TreeNode)
	helper = func(node *TreeNode) {
		if node == nil {
			return
		}
		helper(node.Left)

		if last == math.MaxInt {
			last = node.Val
		}else {
			if 	int(math.Abs(float64(node.Val - last))) < ans {
				ans = int(math.Abs(float64(node.Val - last)))
			}
			last = node.Val
		}

		helper(node.Right)
	}

	helper(root)
	return ans
}

// 501. 二叉搜索树中的众数
// https://leetcode.cn/problems/find-mode-in-binary-search-tree/
func findMode(root *TreeNode) []int {
	ans := make([]int, 0)
	curNum, maxNum, base := 1, 0, math.MinInt64
	var helper func(node *TreeNode)
	helper = func(node *TreeNode) {
		if node == nil {	// 处理最后的元素
			return
		}
		helper(node.Left)

		// 中序遍历处理值结果
		if base == node.Val {
			curNum++
		}else {
			base = node.Val
			curNum = 1
		}
		if curNum == maxNum {
			ans = append(ans, base)
		}else if curNum > maxNum {
			maxNum = curNum
			ans = []int{base}
		}

		helper(node.Right)
	}
	helper(root)

	return ans
}

// 701. 二叉搜索树中的插入操作
// https://leetcode.cn/problems/insert-into-a-binary-search-tree/
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	if root.Val > val {
		root.Left = insertIntoBST(root.Left, val)
	}else {
		root.Right = insertIntoBST(root.Right, val)
	}
	return root
}

// 450. 删除二叉搜索树中的节点
// https://leetcode.cn/problems/delete-node-in-a-bst/
// 第一种情况：没找到删除的节点，遍历到空节点直接返回了
//找到删除的节点
//第二种情况：左右孩子都为空（叶子节点），直接删除节点， 返回NULL为根节点
//第三种情况：删除节点的左孩子为空，右孩子不为空，删除节点，右孩子补位，返回右孩子为根节点
//第四种情况：删除节点的右孩子为空，左孩子不为空，删除节点，左孩子补位，返回左孩子为根节点
//第五种情况：左右孩子节点都不为空，则将删除节点的左子树头结点（左孩子）放到删除节点的右子树的最左面节点的左孩子上，返回删除节点右孩子为新的根节点。
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {	//1.
		return nil
	}

	if root.Val == key {
		//2.
		if root.Left == nil && root.Right == nil {
			return nil
		}
		//3.
		if root.Left == nil {
			return root.Right
		}
		//4.
		if root.Right == nil {
			return root.Left
		}
		//5.
		if root.Left != nil && root.Right != nil {
			cur := root.Right
			for cur.Left != nil {	// 到达最左边
				cur = cur.Left
			}
			cur.Left = root.Left
			root = root.Right	// 替换删除的节点
			return root
		}
	}

	if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	}
	if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	}
	return root
}

// 669. 修剪二叉搜索树
// https://leetcode.cn/problems/trim-a-binary-search-tree/
func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {	// 终止条件
		return nil
	}
	
	// 剪枝操作
	if root.Val < low {
		return trimBST(root.Right, low, high)
	}
	if root.Val > high {
		return trimBST(root.Left, low, high)
	}
	
	root.Left = trimBST(root.Left, low, high)
	root.Right = trimBST(root.Right, low, high)
	return root
}