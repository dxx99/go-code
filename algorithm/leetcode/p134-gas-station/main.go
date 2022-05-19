package main

import "fmt"

// p134 加油站
// 在一条环路上有 n 个加油站，其中第 i 个加油站有汽油 gas[i] 升。
//
//你有一辆油箱容量无限的的汽车，从第 i 个加油站开往第 i+1 个加油站需要消耗汽油 cost[i] 升。你从其中的一个加油站出发，开始时油箱为空。
//
//给定两个整数数组 gas 和 cost ，如果你可以绕环路行驶一周，则返回出发时加油站的编号，否则返回 -1 。如果存在解，则 保证 它是 唯一 的。
//
// 
//
//示例 1:
//
//输入: gas = [1,2,3,4,5], cost = [3,4,5,1,2]
//输出: 3
//解释:
//从 3 号加油站(索引为 3 处)出发，可获得 4 升汽油。此时油箱有 = 0 + 4 = 4 升汽油
//开往 4 号加油站，此时油箱有 4 - 1 + 5 = 8 升汽油
//开往 0 号加油站，此时油箱有 8 - 2 + 1 = 7 升汽油
//开往 1 号加油站，此时油箱有 7 - 3 + 2 = 6 升汽油
//开往 2 号加油站，此时油箱有 6 - 4 + 3 = 5 升汽油
//开往 3 号加油站，你需要消耗 5 升汽油，正好足够你返回到 3 号加油站。
//因此，3 可为起始索引。
//示例 2:
//
//输入: gas = [2,3,4], cost = [3,4,3]
//输出: -1
//解释:
//你不能从 0 号或 1 号加油站出发，因为没有足够的汽油可以让你行驶到下一个加油站。
//我们从 2 号加油站出发，可以获得 4 升汽油。 此时油箱有 = 0 + 4 = 4 升汽油
//开往 0 号加油站，此时油箱有 4 - 3 + 2 = 3 升汽油
//开往 1 号加油站，此时油箱有 3 - 3 + 3 = 3 升汽油
//你无法返回 2 号加油站，因为返程需要消耗 4 升汽油，但是你的油箱只有 3 升汽油。
//因此，无论怎样，你都不可能绕环路行驶一周。
// 
//
//提示:
//
//gas.length == n
//cost.length == n
//1 <= n <= 105
//0 <= gas[i], cost[i] <= 104
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/gas-station
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	fmt.Println(canCompleteCircuit([]int{1,2,3,4,5}, []int{3,4,5,1,2}))
	fmt.Println(canCompleteCircuit([]int{2,3,4}, []int{3,4,3}))

	fmt.Println(canCompleteCircuitV2([]int{1,2,3,4,5}, []int{3,4,5,1,2}))
	fmt.Println(canCompleteCircuitV2([]int{2,3,4}, []int{3,4,3}))
}

// 贪心算法
func canCompleteCircuitV2(gas []int, cost []int) int {
	curSum, totalSum := 0, 0
	start := 0
	for i := 0; i < len(gas); i++ {
		curSum += gas[i] - cost[i]
		totalSum += gas[i] - cost[i]

		if curSum < 0 {
			start = i + 1
			curSum = 0
		}
	}

	if totalSum < 0 {	// 爬完一圈总和小于零，说明一定跑不到终点
		return -1
	}

	return start
}


// 保利求解，通过不了
func canCompleteCircuit(gas []int, cost []int) int {
	for i := 0; i < len(gas); i++ {
		if gas[i] - cost[i] <= 0 {
			continue
		}
		// 起始位置
		start, end := i, i + len(gas)
		tmp := 0
		for j := start; j <= end; j++ {
			k := j % len(gas)
			tmp += gas[k] - cost[k]
			if tmp < 0 {
				break
			}
		}

		// 说明存在这样的结果，可以返回
		if tmp >= 0 {
			return start
		}
	}
	return -1
}
