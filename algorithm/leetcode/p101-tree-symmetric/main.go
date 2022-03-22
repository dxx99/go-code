package main

import "fmt"

// 101. 对称二叉树
// 给你一个二叉树的根节点 root ， 检查它是否轴对称。
//
// 
//
//示例 1： https://assets.leetcode.com/uploads/2021/02/19/symtree1.jpg
//输入：root = [1,2,2,3,4,4,3]
//输出：true
//
//示例 2：https://assets.leetcode.com/uploads/2021/02/19/symtree2.jpg
//输入：root = [1,2,2,null,3,null,3]
//输出：false
// 
//
//提示：
//树中节点数目在范围 [1, 1000] 内
//-100 <= Node.val <= 100
// 
//
//进阶：你可以运用递归和迭代两种方法解决这个问题吗？
//
//
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/symmetric-tree
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	a := []int{1,2,3,4,5,6,7}
	fmt.Println(a[2:])

}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// 递归实现, 左右相等
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return helper(root.Left, root.Right)
}

func helper(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	// 只要有一个错误就退出
	if left == nil || right == nil || left.Val != right.Val  {
		return false
	}
	return helper(left.Left, right.Right) && helper(left.Right, right.Left)
}


// 迭代逻辑解法
// 利用一个队列将递归转换成迭代，当队列为空，或者检测到不对称时就结束
func isSymmetricV2(root *TreeNode) bool {
	queue := make([]*TreeNode, 0)

	if root == nil {
		return true
	}
	
	queue = append(queue, root, root)
	for len(queue) != 0 {
		//取数据比较
		left, right := queue[0], queue[1]
		queue = queue[2:]

		// 这里不是结束条件，有可能队列还有其他元素没有处理完
		if left == nil && right == nil {
			continue
		}

		if left == nil || right == nil || left.Val != right.Val {
			return false
		}

		// 添加到队列
		queue = append(queue, left.Left, right.Right)
		queue = append(queue, left.Right, right.Left)
	}
	return true
}
