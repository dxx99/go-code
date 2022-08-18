package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	
	//fmt.Println(numberOfRounds("12:01", "12:02"))
	fmt.Println(numberOfRounds("09:31", "10:14"))
}

// 1525. 字符串的好分割数目
// https://leetcode.cn/problems/number-of-good-ways-to-split-a-string/
func numSplits(s string) int {
	total := 0
	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		if _, ok :=m[s[i]]; !ok {
			total++
		}
		m[s[i]]++
	}
	ans := 0

	left := make(map[byte]int)
	leftTotal := 0
	for i := 0; i < len(s); i++ {
		if _, ok :=left[s[i]]; !ok {
			leftTotal++
		}
		left[s[i]]++

		m[s[i]]--
		if m[s[i]] == 0 {
			total--
		}

		if total == leftTotal {
			 ans++
		}
	}

	return ans
}


// 1968. 构造元素不等于两相邻元素平均值的数组
// https://leetcode.cn/problems/array-with-elements-not-equal-to-average-of-neighbors/
func rearrangeArray(nums []int) []int {
	// 使用随机算法
	rand.Seed(time.Now().Unix())

outer:
	for {
		rand.Shuffle(len(nums), func(i, j int) {
			nums[i], nums[j] = nums[j], nums[i]
		})
		for i := 1; i < len(nums)-1; i++ {
			if nums[i]*2 == nums[i-1]+nums[i+1] {
				goto outer
			}
		}

		return nums
	}
}

// 1904. 你完成的完整对局数
// https://leetcode.cn/problems/the-number-of-full-rounds-you-have-played/
func numberOfRounds(loginTime string, logoutTime string) int {
	var h1, h2, m1 ,m2 int
	_, _ = fmt.Sscanf(loginTime, "%d:%d", &h1, &m1)
	_, _ = fmt.Sscanf(logoutTime, "%d:%d", &h2, &m2)
	if loginTime > logoutTime {	// 玩了通宵
		h2 += 24
	}

	s, t := h1*60+m1, h2*60+m2
	//fmt.Println(s, t)
	return (t-t%15-s)/15
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
	return nil

}



