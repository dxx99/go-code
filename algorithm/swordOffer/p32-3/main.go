package main

// 请实现一个函数按照之字形顺序打印二叉树，即第一行按照从左到右的顺序打印，第二层按照从右到左的顺序打印，第三行再按照从左到右的顺序打印，其他行以此类推。
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
//  [20,9],
//  [15,7]
//]
// 
//
//提示：
//
//节点总数 <= 1000
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/cong-shang-dao-xia-da-yin-er-cha-shu-iii-lcof
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {

}

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
	k := 0
	queue = append(queue, root)
	for len(queue) > 0 {
		k++
		item := make([]int, 0)
		tmp :=make([]*TreeNode, 0)
		for _, node := range queue {
			// 这里也可以使用数组的反转，这样内存的消耗低一点
			if k % 2 == 1 {
				item = append(item, node.Val)
			}else{
				item = append([]int{node.Val}, item...)

			}
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


