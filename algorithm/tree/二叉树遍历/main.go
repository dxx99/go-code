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
		ans = append(ans, node.Val)		// 注意：根节点的位置
		traversal(node.Left)
		traversal(node.Right)
	}

	traversal(root)
	return ans
}

// 中序遍历
func midOrderTraversal(root *TreeNode) []int {
	ans := make([]int, 0)

	var traversal func(*TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		traversal(node.Left)
		ans = append(ans, node.Val)	// 注意：根节点的位置
		traversal(node.Right)
	}

	traversal(root)
	return ans
}


// 后序遍历
func endOrderTraversal(root *TreeNode) []int {
	ans := make([]int, 0)

	var traversal func(*TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		traversal(node.Left)
		traversal(node.Right)
		ans = append(ans, node.Val)	// 注意：根节点的位置
	}

	traversal(root)
	return ans
}


// 迭代法 [前序遍历]
// 先将根节点放入栈中，然后右节点放入栈中，再将左节点放入栈中
// 先右再左的原因：出栈的时候就能中左右的循序
func preOrderTraversalV2(root *TreeNode) []int {
	ans := make([]int, 0)

	stack := make([]*TreeNode, 0)
	if root == nil {
		return ans
	}

	stack = append(stack, root)
	for len(stack) != 0 {
		// 出栈
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// 存储值
		ans = append(ans, node.Val)

		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}

	return ans
}

// 迭代法 [中序遍历]
// 需要通过指针来配合使用
func midOrderTraversalV2(root *TreeNode) []int {
	ans := make([]int, 0)
	stack := make([]*TreeNode, 0)

	cur := root
	for cur != nil || len(stack) != 0 {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		} else {	// 这种时候已经到达了左边的最底层

			// 弹出元素
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			ans = append(ans, cur.Val)

			cur = cur.Right
		}
	}

	return ans
}

// 迭代法 [后续遍历]
// 先进行【中右左】访问，然后再反转【左右中】
func endOrderTraversalV2(root *TreeNode) []int {
	ans := make([]int, 0)
	stack := make([]*TreeNode, 0)

	stack = append(stack, root)
	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		ans = append(ans, node.Val)
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}

	// 反转结果
	reverse := func(nums []int) {
		left, right := 0, len(nums)-1
		for left < right {
			nums[left], nums[right] = nums[right], nums[left]
			left++
			right--
		}
	}
	reverse(ans)
	return ans
}