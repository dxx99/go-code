package main

import "fmt"

// p77 组合
// 给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。
//
//你可以按 任何顺序 返回答案。
//
// 
//
//示例 1：
//
//输入：n = 4, k = 2
//输出：
//[
//  [2,4],
//  [3,4],
//  [2,3],
//  [1,2],
//  [1,3],
//  [1,4],
//]
//示例 2：
//
//输入：n = 1, k = 1
//输出：[[1]]
// 
//
//提示：
//
//1 <= n <= 20
//1 <= k <= n
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/combinations
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	fmt.Println(combine(4,2))
	fmt.Println(combineV2(4,2))
}


// 使用闭包优化写法
func combineV2(n int, k int) [][]int {
	rds := make([][]int, 0)	 // 存放符合条件结果的集合
	track := make([]int, 0)	 // 存放符合条件单一结果

	var backtrack func(n, k, start int)
	backtrack = func(n, k, start int) {
		// 终止条件
		if len(track) == k {
			tmp := make([]int, k)
			copy(tmp, track)
			rds = append(rds, tmp)
			return
		}

		// 剪枝优化：可以剪枝的地方在递归中每一层的for循环所选择的起始位置
		// 如果for循环选择的开始位置之后的元素个数 已经不足我们需要的元素个数，则没有必要搜索了
		// 优化过程：
		//		1. 已选择的元素个数   len(track)
		//		2. 还需要的元素个数	k - len(track)
		//		3. 在集合中至多需要从该起始位置开始遍历  n - (k - len(track)) + 1
		//			- eg.  n = 4, k = 3, len(track) = 0,  最大开始的位置为 2
		//
		// 遍历
		for i := start; i <= n - (k - len(track)) + 1; i++ {
			track = append(track, i)

			// 递归
			backtrack(n, k, i+1)

			// 回溯, 清除上一个值
			track = track[:len(track)-1]
		}
	}

	backtrack(n, k, 1)
	return rds
}

// 需要定义一个全局变量存储值
var (
	res [][]int
)

func combine(n int, k int) [][]int {
	// 初始化
	res = make([][]int, 0)

	//特殊情况判断
	if n <= 0 || k <= 0 || k > n {
		return res
	}

	backtracking(n, k, 1, []int{})

	return res
}

func backtracking(n int, k int, start int, track []int) {
	// 终止条件
	if len(track) == k {
		tmp := make([]int, k)
		copy(tmp, track)
		res = append(res, tmp)
		return
	}

	// 遍历
	for i := start; i <= n; i++ {
		// 条件处理
		track = append(track, i)

		// 递归
		backtracking(n, k, i+1, track)

		// 回溯, 撤销条件
		track = track[:len(track)-1]
	}
}
