package main

// 从上到下打印出二叉树的每个节点，同一层的节点按照从左到右的顺序打印。
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
//返回：
//
//[3,9,20,15,7]
// 
//
//提示：
//
//节点总数 <= 1000
//
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/cong-shang-dao-xia-da-yin-er-cha-shu-lcof
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


// 使用辅助队列来进行操作
// 层序遍历
func levelOrder(root *TreeNode) []int {
	queue := make([]*TreeNode, 0)
	ans := make([]int, 0)
	if root == nil {
		return ans
	}

	queue = append(queue, root)
	for len(queue) > 0 {
		// 出列所有元素
		item := queue[0]
		queue = queue[1:]
		ans = append(ans, item.Val)

		// 入列
		if item.Left != nil {
			queue = append(queue, item.Left)
		}
		if item.Right != nil {
			queue = append(queue, item.Right)
		}
	}

	return ans
}


