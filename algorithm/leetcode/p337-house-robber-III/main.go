package main

// p337 打家劫舍III
// 小偷又发现了一个新的可行窃的地区。这个地区只有一个入口，我们称之为 root 。
//
//除了 root 之外，每栋房子有且只有一个“父“房子与之相连。一番侦察之后，聪明的小偷意识到“这个地方的所有房屋的排列类似于一棵二叉树”。 如果 两个直接相连的房子在同一天晚上被打劫 ，房屋将自动报警。
//
//给定二叉树的 root 。返回 在不触动警报的情况下 ，小偷能够盗取的最高金额 。
//
// 
//
//示例 1:
//
//
//
//输入: root = [3,2,3,null,3,null,1]
//输出: 7
//解释: 小偷一晚能够盗取的最高金额 3 + 3 + 1 = 7
//示例 2:
//
//
//
//输入: root = [3,4,5,1,3,null,1]
//输出: 9
//解释: 小偷一晚能够盗取的最高金额 4 + 5 = 9
// 
//
//提示：
//
//树的节点数在 [1, 104] 范围内
//0 <= Node.val <= 104
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/house-robber-iii
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {

}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	dp := robTree(root)
	return max(dp[0], dp[1])
}

func robTree(cur *TreeNode) []int {
	if cur == nil {
		return []int{0, 0}
	}

	// dp数组的定义，dp[0]表示不偷该节点的值，dp[1]表示偷该节点的值

	// 后续遍历
	left := robTree(cur.Left)
	right := robTree(cur.Right)


	// 递推公式： 考虑去偷当前的屋子
	robCur := cur.Val + left[0] + right[0]
	// 不偷当前节点的值
	noRobCur := max(left[0], left[1]) + max(right[0], right[1])

	// 注意顺序，不偷与偷
	return []int{noRobCur, robCur}
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}