package main

import (
	"fmt"
	"sort"
)

// p435 无重叠区间
// 给定一个区间的集合 intervals ，其中 intervals[i] = [starti, endi] 。返回 需要移除区间的最小数量，使剩余区间互不重叠 。
//
// 
//
//示例 1:
//
//输入: intervals = [[1,2],[2,3],[3,4],[1,3]]
//输出: 1
//解释: 移除 [1,3] 后，剩下的区间没有重叠。
//示例 2:
//
//输入: intervals = [ [1,2], [1,2], [1,2] ]
//输出: 2
//解释: 你需要移除两个 [1,2] 来使剩下的区间没有重叠。
//示例 3:
//
//输入: intervals = [ [1,2], [2,3] ]
//输出: 0
//解释: 你不需要移除任何区间，因为它们已经是无重叠的了。
// 
//
//提示:
//
//1 <= intervals.length <= 105
//intervals[i].length == 2
//-5 * 104 <= starti < endi <= 5 * 104
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/non-overlapping-intervals
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	//[[1,100],[11,22],[1,11],[2,12]]
	fmt.Println(eraseOverlapIntervals([][]int{{1,100},{11,22},{1,11},{2,12}}))
	fmt.Println(eraseOverlapIntervals([][]int{{1,2},{2,3},{3,4},{1,3},{2,4}}))
	fmt.Println(eraseOverlapIntervals([][]int{{1,2},{1,2},{1,2},{1,2},{1,2}}))
	fmt.Println(eraseOverlapIntervals([][]int{{1,2},{2,3}}))
}

// 贪心算法
// 思路：题目一看就感觉要排序，究竟是按照右边界排序，还是左边界排序
func eraseOverlapIntervals(intervals [][]int) int {
	ans := 0
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][1] == intervals[j][1] {
			return intervals[i][0] > intervals[j][0]
		}
		return intervals[i][1] < intervals[j][1]
	})

	p := intervals[0][1]	// 记录上次的到达的位置，也就是区间分隔点
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < p {
			ans++
			continue
		}
		p = intervals[i][1]
	}

	return ans
}
