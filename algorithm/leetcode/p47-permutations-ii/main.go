package main

import (
	"fmt"
	"sort"
)

// p47 全排列II
// 给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。
//
// 
//
//示例 1：
//
//输入：nums = [1,1,2]
//输出：
//[[1,1,2],
// [1,2,1],
// [2,1,1]]
//示例 2：
//
//输入：nums = [1,2,3]
//输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
// 
//
//提示：
//
//1 <= nums.length <= 8
//-10 <= nums[i] <= 10
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/permutations-ii
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	fmt.Println(permuteUnique([]int{1,1,2}))
	fmt.Println(permuteUnique([]int{1,2,3}))
}

func permuteUnique(nums []int) [][]int {
	res := make([][]int, 0)
	track := make([]int, 0)
	used := [21]int{}

	// 历史数据
	history := [21]int{}
	for i := 0; i < len(nums); i++ {
		history[nums[i]+10]++
	}

	sort.Ints(nums)

	var backtracking func()
	backtracking = func() {
		if len(nums) == len(track) {
			tmp := make([]int, len(track))
			copy(tmp, track)
			res = append(res, tmp)
			return
		}

		//遍历
		for i := 0; i < len(nums); i++ {
			if i > 0 && nums[i] == nums[i-1] {
				continue
			}

			if used[nums[i]+10] != history[nums[i]+10] {
				used[nums[i]+10]++
				track = append(track, nums[i])
				backtracking()

				// 回溯
				used[nums[i]+10]--
				track = track[:len(track)-1]
			}
		}
	}

	backtracking()
	return res
}
