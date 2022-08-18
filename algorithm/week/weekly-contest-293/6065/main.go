package main

import "fmt"

func main() {
	fmt.Println(largestCombination([]int{16,17,71,62,12,24,14}))
	fmt.Println(largestCombination([]int{8,8}))
}

// dynamicProgramming
// TODO:
func largestCombinationV2(candidates []int) int {
	max := 0
	dp := make([][]int, len(candidates))

	// 初始化
	dp[0] = []int{candidates[0], 1}

	// 递推公式  dp[j] = dp[j-1][1] + 1




	return max
}

// 回溯求解
func largestCombination(candidates []int) int {
	max := 0
	track := make([]int, 0)

	var backtracking func(start int)
	backtracking = func(start int) {
		if len(track) > 0 {
			cur := getByte(track)
			if cur > 0 &&  len(track) > max {
				max = len(track)
			}
		}
		if start > len(candidates) {
			return
		}

		// 遍历
		for i := start; i < len(candidates); i++ {
			track = append(track, candidates[i])

			// 递归
			backtracking(i+1)

			//回溯
			track = track[:len(track)-1]
		}
	}

	backtracking(0)
	return max
}

func getByte(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	res := arr[0]
	for i := 1; i < len(arr); i++ {
		res &= arr[i]
	}
	return res
}
