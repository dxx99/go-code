package main
// 102. 二叉树的层序遍历
// 给你二叉树的根节点 root ，返回其节点值的 层序遍历 。 （即逐层地，从左到右访问所有节点）。
//
// 
//
//示例 1：https://assets.leetcode.com/uploads/2021/02/19/tree1.jpg
//
//
//输入：root = [3,9,20,null,null,15,7]
//输出：[[3],[9,20],[15,7]]
//示例 2：
//
//输入：root = [1]
//输出：[[1]]
//示例 3：
//
//输入：root = []
//输出：[]
// 
//
//提示：
//
//树中节点数目在范围 [0, 2000] 内
//-1000 <= Node.val <= 1000
//
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/binary-tree-level-order-traversal
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {

}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// 广度优先搜索  使用递归和队列
func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	queue := make([]*TreeNode, 0)

	if root == nil {
		return res
	}

	queue = append(queue, root)
	for len(queue) > 0 {
		tmp := make([]int, 0)
		tmpQ := make([]*TreeNode, 0)
		for _, node := range queue {
			tmp = append(tmp, node.Val)
			if node.Left != nil {
				tmpQ = append(tmpQ, node.Left)
			}
			if node.Right != nil {
				tmpQ = append(tmpQ, node.Right)
			}
		}
		res = append(res, tmp)

		//替换
		queue = tmpQ
	}

	return res
}
