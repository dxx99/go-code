package main

import "fmt"

// p491. 递增子序列
// 给你一个整数数组 nums ，找出并返回所有该数组中不同的递增子序列，递增子序列中 至少有两个元素 。你可以按 任意顺序 返回答案。
//
//数组中可能含有重复元素，如出现两个整数相等，也可以视作递增序列的一种特殊情况。
//
// 
//
//示例 1：
//
//输入：nums = [4,6,7,7]
//输出：[[4,6],[4,6,7],[4,6,7,7],[4,7],[4,7,7],[6,7],[6,7,7],[7,7]]
//示例 2：
//
//输入：nums = [4,4,3,2,1]
//输出：[[4,4]]
// 
//
//提示：
//
//1 <= nums.length <= 15
//-100 <= nums[i] <= 100
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/increasing-subsequences
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	//fmt.Println(findSubsequences([]int{100,90,80,70,60,50,60,70,80,90,100}))
	//fmt.Println(findSubsequences([]int{1,2,3,4,5,6,7,8,9,10,1,1,1,1,1}))
	fmt.Println(findSubsequencesV2([]int{4,6,7,7}))
	fmt.Println(findSubsequencesV2([]int{4,4,3,2,1}))
}

// 优化代码
// 使用数组代替hash
// 优化判断条件
func findSubsequencesV2(nums []int) [][]int {
	res := make([][]int, 0)
	track := make([]int, 0)


	var backtracking func(start int)
	backtracking = func(start int) {
		// 终止条件
		if len(track) >= 2 {
			tmp := make([]int, len(track))
			copy(tmp, track)
			res = append(res, tmp)
		}

		// 结束
		if start >= len(nums) {
			return
		}


		// 遍历
		used := [201]int{}
		for i := start; i < len(nums); i++ {
			// 只有两种情况需要跳过
			// 1. 当前元素小于track中的最后一个元素的时候，需要跳过
			// 2. 当前行中已经出现过这个元素了，则不需要再次出现，则需要跳过
			if (len(track) > 0 && track[len(track)-1] > nums[i]) || used[nums[i] + 100] == 1 {
				continue
			}

			track = append(track, nums[i])
			used[nums[i] + 100]++

			// 递归+回溯
			backtracking(i+1)
			track = track[:len(track)-1]
		}
	}

	backtracking(0)
	return res
}


func findSubsequences(nums []int) [][]int {
	res := make([][]int, 0)
	track := make([]int, 0)


	var backtracking func(start int)
	backtracking = func(start int) {
		// 终止条件
		if len(track) >= 2 {
			tmp := make([]int, len(track))
			copy(tmp, track)
			res = append(res, tmp)
		}

		// 结束
		if start >= len(nums) {
			return
		}


		// 遍历
		used := make(map[int]int)
		for i := start; i < len(nums); i++ {
			if _, ok := used[nums[i]]; ok {
				continue
			}
			// 去重
			if i > start && nums[i-1] == nums[i] {
				continue
			}

			if (len(track) > 0 && track[len(track)-1] <= nums[i]) || len(track) == 0 {
				track = append(track, nums[i])

				used[nums[i]]++

				// 递归+回溯
				backtracking(i+1)

				track = track[:len(track)-1]
				
			}
		}
	}

	backtracking(0)
	return res
}
