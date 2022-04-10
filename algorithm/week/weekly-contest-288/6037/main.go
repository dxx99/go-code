package main

import (
	"fmt"
	"sort"
	"strconv"
)

// 6037. 按奇偶性交换后的最大数字 显示英文描述
//通过的用户数0
//尝试过的用户数0
//用户总通过次数0
//用户总提交次数0
//题目难度Easy
//给你一个正整数 num 。你可以交换 num 中 奇偶性 相同的任意两位数字（即，都是奇数或者偶数）。
//
//返回交换 任意 次之后 num 的 最大 可能值。
//
//
//
//示例 1：
//
//输入：num = 1234
//输出：3412
//解释：交换数字 3 和数字 1 ，结果得到 3214 。
//交换数字 2 和数字 4 ，结果得到 3412 。
//注意，可能存在其他交换序列，但是可以证明 3412 是最大可能值。
//注意，不能交换数字 4 和数字 1 ，因为它们奇偶性不同。
//示例 2：
//
//输入：num = 65875
//输出：87655
//解释：交换数字 8 和数字 6 ，结果得到 85675 。
//交换数字 5 和数字 7 ，结果得到 87655 。
//注意，可能存在其他交换序列，但是可以证明 87655 是最大可能值。
//
//
//提示：
//
//1 <= num <= 109
// https://leetcode-cn.com/contest/weekly-contest-288/problems/largest-number-after-digit-swaps-by-parity/
func main() {
	fmt.Println(largestInteger(1234))
	fmt.Println(largestInteger(65875))
}

func largestInteger(num int) int {
	strArr := strconv.Itoa(num)

	// 存放偶数
	hashMap := make(map[int]int, 0)
	oddArr := make([]int, 0)
	evenArr := make([]int, 0)
	for _, ch := range strArr {
		key := int(ch - '0')
		if key % 2 == 0 {
			evenArr = append(evenArr, key)
		} else {
			oddArr = append(oddArr, key)
		}
		hashMap[key]++
	}
	sort.Ints(oddArr)
	sort.Ints(evenArr)

	resArr := make([]int, 0)

	eIndex, oIndex := len(evenArr)-1, len(oddArr)-1
	for _, ch := range strArr {
		key := int(ch - '0')
		if key % 2 == 0 {
			resArr = append(resArr, evenArr[eIndex])
			eIndex--
		}else {
			resArr = append(resArr, oddArr[oIndex])
			oIndex--
		}
	}

	maxNum := 0
	for _, n := range resArr {
		maxNum = maxNum*10 + n
	}

	return maxNum
}
