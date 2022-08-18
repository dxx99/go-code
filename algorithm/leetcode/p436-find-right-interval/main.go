package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findRightInterval([][]int{{1,1},{3,4}}))
	fmt.Println(findRightIntervalV2([][]int{{1,1},{3,4}}))
	//fmt.Println(findRightInterval([][]int{{3,4},{2,3},{1,2}}))
}

// 436. 寻找右区间
// https://leetcode.cn/problems/find-right-interval/
func findRightIntervalV2(intervals [][]int) []int {
	// 	区间i 右侧区间j, 并且 start[j] >= end[i], 且start[j]最小化
	for i := 0; i < len(intervals); i++ {
		intervals[i] = append(intervals[i], i)
	}

	//sort
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})

	ans := make([]int, len(intervals))
	for i := 0; i < len(ans); i++ {
		ans[i] = -1
	}

	for i := 0; i < len(intervals); i++ {
		x := sort.Search(len(intervals), func(k int) bool {
			return intervals[k][0] >= intervals[i][1]
		})
		fmt.Println(intervals, i, x)
		if x < len(intervals) {
			ans[intervals[i][2]] = intervals[x][2]
		}
	}
	return ans
}

func findRightInterval(intervals [][]int) []int {
	// 	区间i 右侧区间j, 并且 start[j] >= end[i], 且start[j]最小化
	for i := 0; i < len(intervals); i++ {
		intervals[i] = append(intervals[i], i)
	}

	//sort
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})

	ans := make([]int, len(intervals))
	for i := 0; i < len(ans); i++ {
		ans[i] = -1
	}
	for i := 0; i < len(intervals); i++ {
		for j := i; j < len(intervals); j++ {	// 自己的左边可以是自己
			if intervals[j][0] >= intervals[i][1] {
				ans[intervals[i][2]] = intervals[j][2]
				break
			}
		}
	}

	return ans
}