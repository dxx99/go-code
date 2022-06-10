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

// 112. 路径总和
// https://leetcode.cn/problems/path-sum/
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	cur := root.Val
	var helper func(node *TreeNode) bool
	helper = func(node *TreeNode) bool {
		if node.Left == nil && node.Right == nil && cur == targetSum {
			return true
		}
		if node.Left != nil {
			cur += node.Left.Val
			leftRes := helper(node.Left)
			if leftRes {
				return true
			}
			cur -= node.Left.Val	//回溯逻辑

		}
		if node.Right != nil {
			cur += node.Right.Val
			rightRes := helper(node.Right)
			if rightRes {
				return true
			}
			cur -= node.Right.Val
		}
		return false
	}
	return helper(root)
}
// 优化递归代码
func hasPathSumV2(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil && targetSum == root.Val {
		return true
	}
	return hasPathSumV2(root.Left, targetSum - root.Val) || hasPathSumV2(root.Right, targetSum-root.Val)
}


// 113. 路径总和 II
// https://leetcode.cn/problems/path-sum-ii/
func pathSum(root *TreeNode, targetSum int) [][]int {
	ans := make([][]int, 0)
	if root == nil {
		return ans
	}

	track := make([]int, 0)
	var helper func(node *TreeNode, sum int)
	helper = func(node *TreeNode, sum int) {
		if node.Left == nil && node.Right == nil && sum == 0 {
			tmp := make([]int, len(track))	// 回溯数组存储，要copy
			copy(tmp, track)
			ans = append(ans, tmp)
			return
		}

		if node.Left != nil {
			track = append(track, node.Left.Val)
			helper(node.Left, sum - node.Left.Val)
			track = track[:len(track)-1]	// 回溯操作
		}

		if node.Right != nil {
			track = append(track, node.Right.Val)
			helper(node.Right, sum - node.Right.Val)
			track = track[:len(track)-1]
		}
	}
	track = append(track, root.Val)
	helper(root, targetSum - root.Val)
	return ans
}
// 代码优化
func pathSumV2(root *TreeNode, targetSum int) [][]int {
	ans := make([][]int, 0)
	track := make([]int, 0)
	var helper func(*TreeNode, int)
	helper = func(node *TreeNode, sum int) {
		if node == nil {
			return
		}

		sum -= node.Val
		track = append(track, node.Val)
		defer func() {track = track[:len(track)-1]}()	// 回溯操作

		if node.Left == nil && node.Right == nil && sum == 0 {
			tmp := make([]int, len(track))	// 回溯数组存储，要copy
			copy(tmp, track)
			ans = append(ans, tmp)
			return
		}

		helper(node.Left, sum)
		helper(node.Right, sum)
	}
	helper(root, targetSum)
	return ans
}

func pathSumV3(root *TreeNode, targetSum int) (ans [][]int) {
	path := []int{}
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, left int) {
		if node == nil {
			return
		}
		left -= node.Val
		path = append(path, node.Val)
		defer func() { path = path[:len(path)-1] }()
		if node.Left == nil && node.Right == nil && left == 0 {
			tmp := make([]int, len(path))	// 回溯数组存储，要copy
			copy(tmp, path)
			ans = append(ans, tmp)
			return
		}
		dfs(node.Left, left)
		dfs(node.Right, left)
	}
	dfs(root, targetSum)
	return
}

