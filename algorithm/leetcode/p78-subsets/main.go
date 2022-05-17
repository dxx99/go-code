package main

import "fmt"

// p78 子集
// 给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
//
//解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。
//
// 
//
//示例 1：
//
//输入：nums = [1,2,3]
//输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
//示例 2：
//
//输入：nums = [0]
//输出：[[],[0]]
// 
//
//提示：
//
//1 <= nums.length <= 10
//-10 <= nums[i] <= 10
//nums 中的所有元素 互不相同
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/subsets
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	fmt.Println(subsets([]int{1,2,3}))
	fmt.Println(subsets([]int{0}))
}

// 组合问题和分隔问题， 收集树上的叶子节点
// 子集问题，查找树上的所有叶子节点，其实也是一种组合问题，因为他的集合是无序的，子集{1,2}与{2,1}是一样的
// 无序，取过的元素不会重复取，写回溯的时候，for就要从start开始，而不能从零开始
// 有序，也就是求排列问题，for循环应该从零开始
// 遍历这个数，把所有的节点都都记录下来，也就这个子集
// 终止条件: start大于数组的元素时就需要终止
func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	track := make([]int, 0)

	var backtracking func(start int)
	backtracking = func(start int) {
		// 终止条件
		if start > len(nums) {
			return
		}

		// 存放结果
		tmp := make([]int, len(track))
		copy(tmp, track)
		res = append(res, tmp)

		// 遍历
		for i := start; i < len(nums); i++ {
			track = append(track, nums[i])

			// 递归
			backtracking(i+1)

			// 回溯
			track = track[:len(track)-1]
		}
	}

	backtracking(0)
	return res
}
