package main

import (
	"fmt"
	"sort"
)

// p56 合并区间
// 以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。
//
// 
//
//示例 1：
//
//输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
//输出：[[1,6],[8,10],[15,18]]
//解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
//示例 2：
//
//输入：intervals = [[1,4],[4,5]]
//输出：[[1,5]]
//解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。
// 
//
//提示：
//
//1 <= intervals.length <= 104
//intervals[i].length == 2
//0 <= starti <= endi <= 104
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/merge-intervals
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	fmt.Println(merge([][]int{{1,3},{2,6},{8,10},{15,18}}))
}

func merge(intervals [][]int) [][]int {
	ans := make([][]int, 0)

	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0]  {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})

	p := intervals[0]
	for i := 1; i < len(intervals); i++ {
		// 对两个集合进行合并
		if intervals[i][0] <= p[1] {
			px, py :=  p[0], p[1]
			if intervals[i][1] > p[1] {
				py = intervals[i][1]
			}
			p = []int{px,py}
			continue
		}

		// 合并不了则需要将p点加入到结果收集集中
		ans = append(ans, p)
		p = intervals[i]
	}

	// 补上最后一个元素
	ans = append(ans, p)
	return ans
}