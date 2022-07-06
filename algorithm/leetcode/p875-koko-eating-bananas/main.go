package main

import "fmt"

func main() {
	fmt.Println(0/10+1)
	fmt.Println(minEatingSpeedV2([]int{3,6,7,11}, 8))
}

// 875. 爱吃香蕉的珂珂
// https://leetcode.cn/problems/koko-eating-bananas/submissions/
func minEatingSpeed(piles []int, h int) int {
	x := 1
	for true {
		usedHour := 0
		for i := 0; i < len(piles); i++ {
			usedHour += piles[i] / x
			if piles[i] % x != 0  {
				usedHour++
			}
		}

		if usedHour <= h {
			break
		}
		x++
	}
	return x
}
// 优化，二分查找
func minEatingSpeedV2(piles []int, h int) int {
	f := func(k int) int {
		usedHour := 0
		for i := 0; i < len(piles); i++ {
			usedHour += piles[i] / k
			if piles[i] % k != 0  {
				usedHour++
			}
		}
		return usedHour
	}

	left, right := 1, int(1e9)
	for left < right {
		mid := int(uint(left+right)>>1)
		if f(mid) > h {
			left = mid+1
		}else {
			right = mid
		}
	}
	return left
}