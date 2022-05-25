package main

// 请实现一个函数，用来判断一棵二叉树是不是对称的。如果一棵二叉树和它的镜像一样，那么它是对称的。
//
//例如，二叉树 [1,2,2,3,4,4,3] 是对称的。
//
//    1
//   / \
//  2   2
// / \ / \
//3  4 4  3
//但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:
//
//    1
//   / \
//  2   2
//   \   \
//   3    3
//
// 
//
//示例 1：
//
//输入：root = [1,2,2,3,4,4,3]
//输出：true
//示例 2：
//
//输入：root = [1,2,2,null,3,null,3]
//输出：false
// 
//
//限制：
//
//0 <= 节点个数 <= 1000
//
//注意：本题与主站 101 题相同：https://leetcode-cn.com/problems/symmetric-tree/
//
//
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/dui-cheng-de-er-cha-shu-lcof
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

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return helper(root.Left, root.Right)
}

func helper(node1 *TreeNode, node2 *TreeNode) bool {
	if node1 == nil && node2 == nil {
		return true
	}

	if node1 == nil || node2 == nil || node1.Val != node2.Val {
		return false
	}

	// 这里可以递归在前，回溯在后面
	return helper(node1.Left, node2.Right) && helper(node1.Right, node2.Left)
}