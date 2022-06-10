package main

import (
	"container/list"
	"fmt"
	"math"
	"strings"
)

func main() {

}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// 226. 翻转二叉树
// https://leetcode.cn/problems/invert-binary-tree/
// https://www.programmercarl.com/0226.%E7%BF%BB%E8%BD%AC%E4%BA%8C%E5%8F%89%E6%A0%91.html
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := invertTree(root.Left)
	right := invertTree(root.Right)
	root.Left, root.Right = right, left // dfs 后序遍历
	return root
}

// 101.对称二叉树
// https://leetcode.cn/problems/symmetric-tree/
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var helper func(left *TreeNode, right *TreeNode) bool
	helper = func(left *TreeNode, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		}
		if (left == nil && right != nil) || (left != nil && right == nil) {
			return false
		}
		return left.Val == right.Val && helper(left.Left, right.Right) && helper(left.Right, right.Left)	// dfs 前序
	}

	return  helper(root.Left, root.Right)
}

// 222. 完全二叉树的节点个数
// https://leetcode.cn/problems/count-complete-tree-nodes/
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l, r := root, root
	hl, hr := 0, 0
	for l != nil {
		l = l.Left
		hl++
	}
	for r != nil {
		r = r.Right
		hr++
	}

	// 当前节点的子树为满二叉树
	if hl == hr {
		return int(math.Pow(float64(2), float64(hl))) - 1
	}
	return countNodes(root.Left) + countNodes(root.Right) + 1	// dfs 后序
}

// 110. 平衡二叉树
// https://leetcode.cn/problems/balanced-binary-tree/
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	// 求高度
	var helper func(node *TreeNode) int	// 传入节点得到高度
	helper = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := helper(node.Left)
		right := helper(node.Right)

		if left > right {	// 后序遍历
			return left+1
		}
		return right+1
	}

	//前序遍历, 可以提前退出
	hl, hr := helper(root.Left), helper(root.Right)
	if math.Abs(float64(hl-hr)) > 1 {
		return false
	}

	return isBalanced(root.Left) && isBalanced(root.Right)
}

// 257. 二叉树的所有路径
// https://leetcode.cn/problems/binary-tree-paths/
func binaryTreePaths(root *TreeNode) []string {
	ans := make([]string, 0)
	if root == nil {
		return ans
	}

	track := make([]string, 0)
	var helper func(node *TreeNode)
	helper = func(node *TreeNode) {
		if node == nil {
			return
		}

		// 这个时候到达叶子节点，可以退出
		if node.Left == nil && node.Right == nil {
			track = append(track, fmt.Sprintf("%d", node.Val))
			ans = append(ans, strings.Join(track, "->"))
			track = track[:len(track)-1]
			return
		}

		track = append(track, fmt.Sprintf("%d", node.Val))	//这里是前序遍历的位置
		helper(node.Left)
		helper(node.Right)
		track = track[:len(track)-1]	// 后序遍历的位置
	}
	
	helper(root)
	return ans
}

// 404.左叶子节点之和
// https://leetcode.cn/problems/sum-of-left-leaves/
func sumOfLeftLeaves(root *TreeNode) int {
	ans := 0
	var helper func(node *TreeNode)
	helper = func(node *TreeNode) {
		if node == nil {
			return
		}
		// 需要判断当前左节点是不是左叶子节点
		if node.Left != nil && node.Left.Left == nil && node.Left.Right == nil {
			ans += node.Left.Val
		}
		if node.Left != nil {
			helper(node.Left)
		}
		if node.Right != nil {
			helper(node.Right)
		}
	}

	helper(root)
	return ans
}

// 513. 找树左下角的值
// https://leetcode.cn/problems/find-bottom-left-tree-value/
func findBottomLeftValue(root *TreeNode) int {
	ans := 0
	if root == nil {
		return ans
	}

	queue := list.New()
	queue.PushBack(root)
	for queue.Len() != 0 {
		size := queue.Len()
		for i := 0; i < size; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if i == 0 {
				ans = node.Val
			}
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
	}
	return ans
}
// 递归法
func findBottomLeftValueV2(root *TreeNode) int {
	ans := 0
	if root == nil {
		return ans
	}

	maxHeight := math.MinInt
	curHeight := 0
	var helper func(node *TreeNode)
	helper = func(node *TreeNode) {
		if node.Left == nil && node.Right == nil {
			if curHeight > maxHeight {	// 记录第一次最大高度的值
				maxHeight = curHeight
				ans = node.Val
			}
		}

		if node.Left != nil {
			curHeight++
			helper(node.Left)
			curHeight--		// 回溯逻辑
		}
		if node.Right != nil {
			curHeight++
			helper(node.Right)
			curHeight--
		}
	}
	helper(root)
	return ans
}

