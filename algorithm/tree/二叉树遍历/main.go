package main


func main() {

}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


// 前序遍历
func preOrderTraversal(root *TreeNode) []int {
	ans := make([]int, 0)

	var traversal func(*TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		ans = append(ans, node.Val)
		traversal(node.Left)
		traversal(node.Right)
	}

	traversal(root)
	return ans
}

// 中序遍历


// 后序遍历
