package main

import (
	"fmt"
	"sort"
)

// p40 组合总和
// 给定一个候选人编号的集合 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
//
//candidates 中的每个数字在每个组合中只能使用 一次 。
//
//注意：解集不能包含重复的组合。 
//
// 
//
//示例 1:
//
//输入: candidates = [10,1,2,7,6,1,5], target = 8,
//输出:
//[
//[1,1,6],
//[1,2,5],
//[1,7],
//[2,6]
//]
//示例 2:
//
//输入: candidates = [2,5,2,1,2], target = 5,
//输出:
//[
//[1,2,2],
//[5]
//]
// 
//
//提示:
//
//1 <= candidates.length <= 100
//1 <= candidates[i] <= 50
//1 <= target <= 30
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/combination-sum-ii
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	fmt.Println(combinationSum2([]int{10,1,2,7,6,1,5}, 8))
	fmt.Println(combinationSum2V2([]int{10,1,2,7,6,1,5}, 8))
	fmt.Println(combinationSum2([]int{2,5,2,1,2}, 5))
	fmt.Println(combinationSum2V2([]int{2,5,2,1,2}, 5))
}

func combinationSum2V2(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	track := make([]int, 0)
	total := 0

	sort.Ints(candidates)	// 排序之后的数据方便去重处理
	var backtracking func(start int)
	backtracking = func(start int) {
		if total > target {
			return
		}
		// 终止条件
		if total == target {
			tmp := make([]int, len(track))
			copy(tmp, track)
			res = append(res, tmp)
			return
		}

		// 遍历
		for i := start; i < len(candidates); i++ {
			if i > start && candidates[i] == candidates[i-1] {	// 只有同一层重复的才需要去重，关键点！！！
				continue
			}
			track = append(track, candidates[i])
			total += candidates[i]
			//递归
			backtracking(i+1)
			//回溯
			track = track[:len(track)-1]
			total -= candidates[i]
		}
	}

	backtracking(0)
	return res
}

// 使用used去重
func combinationSum2(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	track := make([]int, 0)
	total := 0

	used := make(map[int]bool, 0)

	sort.Ints(candidates)	// 排序之后的数据方便去重处理
	var backtracking func(start int)
	backtracking = func(start int) {
		if total > target {
			return
		}
		// 终止条件
		if total == target {
			tmp := make([]int, len(track))
			copy(tmp, track)
			res = append(res, tmp)
			return
		}

		// 遍历
		for i := start; i < len(candidates); i++ {
			if i > 0 && candidates[i] == candidates[i-1] && used[i-1] == false {	// 去重比较关键
				continue
			}
			track = append(track, candidates[i])
			total += candidates[i]
			used[i] = true
			//递归
			backtracking(i+1)
			//回溯
			track = track[:len(track)-1]
			total -= candidates[i]
			used[i] = false
		}
	}

	backtracking(0)
	return res
}
