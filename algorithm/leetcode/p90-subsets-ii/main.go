package main

import (
	"fmt"
	"sort"
)

// p90 子集II
// 给你一个整数数组 nums ，其中可能包含重复元素，请你返回该数组所有可能的子集（幂集）。
//
//解集 不能 包含重复的子集。返回的解集中，子集可以按 任意顺序 排列。
//
// 
//
//示例 1：
//
//输入：nums = [1,2,2]
//输出：[[],[1],[1,2],[1,2,2],[2],[2,2]]
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
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/subsets-ii
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	fmt.Println(subsetsWithDup([]int{4,4,4,1,4}))
	//fmt.Println(subsetsWithDup([]int{1,2,2}))
	//fmt.Println(subsetsWithDup([]int{0}))
}

// 去重问题，先排序，再排除右边的元素
func subsetsWithDup(nums []int) [][]int {
	res := make([][]int, 0)
	track := make([]int, 0)
	sort.Ints(nums)
	var backtracking func(start int)
	backtracking = func(start int) {
		// 终止条件
		if start > len(nums) {
			return
		}

		// 写入数据
		tmp := make([]int, len(track))
		copy(tmp, track)
		res = append(res, tmp)

		// 遍历
		for i := start; i < len(nums); i++ {
			// 除重逻辑
			if i > start && nums[i] == nums[i-1] {
				continue
			}

			track = append(track, nums[i])
			backtracking(i+1)

			// 回溯
			track = track[:len(track)-1]
		}
	}

	backtracking(0)
	return res
}
