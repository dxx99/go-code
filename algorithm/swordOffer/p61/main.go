package main

import "fmt"

// 从若干副扑克牌中随机抽 5 张牌，判断是不是一个顺子，即这5张牌是不是连续的。2～10为数字本身，A为1，J为11，Q为12，K为13，而大、小王为 0 ，可以看成任意数字。A 不能视为 14。
//
// 
//
//示例 1:
//
//输入: [1,2,3,4,5]
//输出: True
// 
//
//示例 2:
//
//输入: [0,0,1,2,5]
//输出: True
// 
//
//限制：
//
//数组长度为 5 
//
//数组的数取值为 [0, 13] .
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/bu-ke-pai-zhong-de-shun-zi-lcof
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	fmt.Println(isStraight([]int{0,2,3,4,6}))
}

// 保持有序的条件
// 1. 无重复
// 2. 最大值-最小值 < 5 其中要跳过大小王
func isStraight(nums []int) bool {
	// 重复问题
	hash := make(map[int]bool)
	max, min := 0, 14
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			continue
		}

		if nums[i] > max {
			max = nums[i]
		}
		if nums[i] < min {
			min = nums[i]
		}

		if _, ok := hash[nums[i]]; ok {
			return false
		}

		hash[nums[i]] = true
	}

	return max - min < 5
}
