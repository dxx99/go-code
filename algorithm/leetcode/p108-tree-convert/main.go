package main

import "fmt"

// 108. 将有序数组转换为二叉搜索树
// 给你一个整数数组 nums ，其中元素已经按 升序 排列，请你将其转换为一棵 高度平衡 二叉搜索树。
//
//高度平衡 二叉树是一棵满足「每个节点的左右两个子树的高度差的绝对值不超过 1 」的二叉树。
//
// 
//
//示例 1：https://assets.leetcode.com/uploads/2021/02/18/btree1.jpg
//
//
//输入：nums = [-10,-3,0,5,9]
//输出：[0,-3,9,-10,null,5]
//解释：[0,-10,5,null,-3,null,9] 也将被视为正确答案：https://assets.leetcode.com/uploads/2021/02/18/btree2.jpg
//
//示例 2： https://assets.leetcode.com/uploads/2021/02/18/btree.jpg
//
//
//输入：nums = [1,3]
//输出：[3,1]
//解释：[1,null,3] 和 [3,1] 都是高度平衡二叉搜索树。
// 
//
//提示：
//
//1 <= nums.length <= 104
//-104 <= nums[i] <= 104
//nums 按 严格递增 顺序排列
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/convert-sorted-array-to-binary-search-tree
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	m := (3+4)/2
	fmt.Println(m)
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// 找到中间节点，然后建立二叉树
func sortedArrayToBST(nums []int) *TreeNode {
	return helper(nums, 0, len(nums) - 1)
}

func helper(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	mid := (left + right)/2		// 选择小的，也就是左边的为根节点， 还可以选择右边的为根节点
	root := &TreeNode{Val: nums[mid]}
	root.Left = helper(nums, left, mid-1)
	root.Right = helper(nums, mid+1, right)
	return root
}
