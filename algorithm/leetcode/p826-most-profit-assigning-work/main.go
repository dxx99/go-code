package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maxProfitAssignment([]int{2,4,6,8,10}, []int{10,20,30,40,50}, []int{4,5,6,7}))
}

// 826. 安排工作以达到最大收益
// https://leetcode.cn/problems/most-profit-assigning-work/
func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	// 工人分配到岗位，最大的利润
	// 每个工作最多只能完成一次任务，并且能力值必须大于困难度，同一个任务可以安排多个工作去完成
	pd := make([][]int, len(profit))	// 利润困难表
	for i := 0; i < len(pd); i++ {
		pd[i] = []int{profit[i], difficulty[i]}
	}

	// 排序
	sort.Slice(pd, func(i, j int) bool {
		return pd[i][0] > pd[j][0]
	})

	// 倒序能力值
	sort.Slice(worker, func(i, j int) bool {
		return worker[i] > worker[j]
	})

	// 最大利润
	ans := 0

	wk := 0
	i := 0
	for i < len(pd) && wk < len(worker) {
		if pd[i][1] <= worker[wk] {
			ans += pd[i][0]
			wk++
			continue
		}
		i++
	}
	return ans
}
