package main

import (
	"container/list"
	"math"
)

func main() {

}


type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// 102.二叉树的层序遍历
// https://leetcode.cn/problems/binary-tree-level-order-traversal/
func levelOrder(root *TreeNode) [][]int {
	ans := make([][]int, 0)
	queue := list.New()
	if root == nil {
		return ans
	}

	queue.PushBack(root)
	for queue.Len() != 0 {
		size := queue.Len()
		tmp := make([]int, 0)
		for i := 0; i < size; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			tmp = append(tmp, node.Val)
		}
		ans = append(ans, tmp)
	}

	return ans
}

// 107.二叉树的层序遍历II
// https://leetcode.cn/problems/binary-tree-level-order-traversal-ii/
func levelOrderBottom(root *TreeNode) [][]int {
	ans := make([][]int, 0)
	queue := list.New()
	if root == nil {
		return ans
	}

	queue.PushBack(root)
	for queue.Len() != 0 {
		size := queue.Len()
		tmp := make([]int, 0)
		for i := 0; i < size; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			tmp = append(tmp, node.Val)
		}
		ans = append(ans, tmp)
	}

	// 反转数组
	reverse := func(nums [][]int) {
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

// 199.二叉树的右视图
// https://leetcode.cn/problems/binary-tree-right-side-view/
func rightSideView(root *TreeNode) []int {
	ans := make([]int, 0)
	queue := make([]*TreeNode, 0)	// 切换数组的写法
	if root == nil {
		return ans
	}

	queue = append(queue, root)
	for len(queue) != 0 {
		newQueue := make([]*TreeNode, 0)
		for i := 0; i < len(queue); i++ {
			node := queue[i]
			if node.Right != nil {
				newQueue = append(newQueue, node.Right)
			}
			if node.Left != nil {
				newQueue = append(newQueue, node.Left)
			}
			if i == 0 {
				ans = append(ans, node.Val)
			}
		}
		queue = newQueue
	}
	return ans
}

// 637.二叉树的层平均值
// https://leetcode.cn/problems/average-of-levels-in-binary-tree/
func averageOfLevels(root *TreeNode) []float64 {
	ans := make([]float64, 0)
	queue := list.New()
	if root == nil {
		return ans
	}

	queue.PushBack(root)
	for queue.Len() != 0 {
		size := queue.Len()
		total := 0	// 保存每一层的和
		for i := 0; i < size; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			total += node.Val
		}
		ans = append(ans, float64(total)/float64(size))
	}

	return ans
}



// Node N叉树定义
type Node struct {
	Val int
	Children []*Node
}

// N叉树的层序遍历
// https://leetcode.cn/problems/n-ary-tree-level-order-traversal/
func levelOrderNTree(root *Node) [][]int {
	ans := make([][]int, 0)
	queue := make([]*Node, 0)

	if root == nil {
		return ans
	}

	queue = append(queue, root)
	for len(queue) != 0 {
		tmp := make([]int, 0)
		newQueue := make([]*Node, 0)
		for i := 0; i < len(queue); i++ {
			node := queue[i]
			tmp = append(tmp, node.Val)
			newQueue = append(newQueue, node.Children...)
		}
		ans = append(ans, tmp)
		queue = newQueue
	}

	return ans
}


// 515.在每个树行中找最大值
// https://leetcode.cn/problems/find-largest-value-in-each-tree-row/
func largestValues(root *TreeNode) []int {
	ans := make([]int, 0)
	queue := list.New()
	if root == nil {
		return ans
	}

	queue.PushBack(root)
	for queue.Len() != 0 {
		size := queue.Len()
		max := math.MinInt	// 保存每一层的和
		for i := 0; i < size; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			if node.Val > max {
				max = node.Val
			}
		}
		ans = append(ans, max)
	}

	return ans
}

type NextNode struct {
	Val int
	Left *NextNode
	Right *NextNode
	Next *NextNode
}

// 116.填充每个节点的下一个右侧节点指针
// https://leetcode.cn/problems/populating-next-right-pointers-in-each-node/
// 117.填充每个节点的下一个右侧节点指针II
// https://leetcode.cn/problems/populating-next-right-pointers-in-each-node-ii/
func connect(root *NextNode) *NextNode {
	queue := list.New()
	if root == nil {
		return root
	}

	queue.PushBack(root)
	for queue.Len() != 0 {
		size := queue.Len()
		var head *NextNode
		for i := 0; i < size; i++ {
			node := queue.Remove(queue.Front()).(*NextNode)
			if head != nil {
				head.Next = node
				head = head.Next
			}else {
				head = node
			}

			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
	}
	return root
}


// 104.二叉树的最大深度
// https://leetcode.cn/problems/maximum-depth-of-binary-tree/
func maxDepth(root *TreeNode) int {
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
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		ans++
	}

	return ans
}

// 111.二叉树的最小深度
// https://leetcode.cn/problems/minimum-depth-of-binary-tree/
func minDepth(root *TreeNode) int {
	ans := 0
	if root == nil {
		return ans
	}
	queue := list.New()
	queue.PushBack(root)

	for queue.Len() != 0 {
		ans++
		size := queue.Len()
		for i := 0; i < size; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Left == nil && node.Right == nil {
				return ans
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
