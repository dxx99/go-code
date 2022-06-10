package main

import "fmt"

//给你二叉树的根节点 root 和一个整数目标和 targetSum ，找出所有 从根节点到叶子节点 路径总和等于给定目标和的路径。
//
//叶子节点 是指没有子节点的节点。
//
// 
//
//示例 1：
//
//
//
//输入：root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
//输出：[[5,4,11,2],[5,8,4,5]]
//示例 2：
//
//
//
//输入：root = [1,2,3], targetSum = 5
//输出：[]
//示例 3：
//
//输入：root = [1,2], targetSum = 0
//输出：[]
// 
//
//提示：
//
//树中节点总数在范围 [0, 5000] 内
//-1000 <= Node.val <= 1000
//-1000 <= targetSum <= 1000
//注意：本题与主站 113 题相同：https://leetcode-cn.com/problems/path-sum-ii/
//
//
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/er-cha-shu-zhong-he-wei-mou-yi-zhi-de-lu-jing-lcof
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	root := NewNode(5)
	root.Left = NewNode(4)
	root.Left.Left = NewNode(11)
	root.Left.Left.Left = NewNode(7)
	root.Left.Left.Right = NewNode(2)
	root.Right = NewNode(8)
	root.Right.Left = NewNode(13)
	root.Right.Right = NewNode(4)
	root.Right.Right.Left = NewNode(5)
	root.Right.Right.Right = NewNode(1)


	fmt.Println(pathSumV2(root, 22))
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

func NewNode(val int) *TreeNode {
	return &TreeNode{
		Val:   val,
		Left:  nil,
		Right: nil,
	}
}

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


// 【 前序遍历】【后序遍历】配合使用
func pathSum(root *TreeNode, target int) [][]int {
	ans := make([][]int, 0)
	track := make([]int, 0)

	if root == nil {
		return ans
	}

	var helper func(node *TreeNode, target int)
	helper = func(node *TreeNode, target int) {
		if node == nil {
			return
		}

		remain := target - node.Val

		// 找到，并且刚好是最后一个元素
		if node.Left == nil && node.Right == nil {
			if remain == 0 {
				track = append(track, node.Val)
				tmp := make([]int, len(track))
				copy(tmp, track)
				ans = append(ans, tmp)

				// 这里就是回溯操作
				track = track[:len(track)-1]
			}
			return
		}


		// 左节点
		track = append(track, node.Val)
		helper(node.Left, remain)
		track = track[:len(track)-1]

		// 右节点
		track = append(track, node.Val)
		helper(node.Right, remain)
		track = track[:len(track)-1]

	}

	// 从根节点开始
	helper(root, target)

	return ans
}


