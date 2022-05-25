package main

// 从上到下按层打印二叉树，同一层的节点按从左到右的顺序打印，每一层打印到一行。
//
// 
//
//例如:
//给定二叉树: [3,9,20,null,null,15,7],
//
//    3
//   / \
//  9  20
//    /  \
//   15   7
//返回其层次遍历结果：
//
//[
//  [3],
//  [9,20],
//  [15,7]
//]
// 
//
//提示：
//
//节点总数 <= 1000
//注意：本题与主站 102 题相同：https://leetcode-cn.com/problems/binary-tree-level-order-traversal/
//
//
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/cong-shang-dao-xia-da-yin-er-cha-shu-ii-lcof
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

func levelOrder(root *TreeNode) [][]int {
	queue := make([]*TreeNode, 0)
	ans := make([][]int, 0)

	if root == nil {
		return ans
	}

	queue = append(queue, root)
	for len(queue) > 0 {
		item := make([]int, 0)
		tmp :=make([]*TreeNode, 0)
		for _, node := range queue {
			item = append(item, node.Val)
			if node.Left != nil {
				tmp = append(tmp, node.Left)
			}
			if node.Right != nil {
				tmp = append(tmp, node.Right)
			}
		}

		queue = tmp
		ans = append(ans, item)
	}

	return ans
}
