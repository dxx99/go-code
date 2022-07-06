package main

import "math"

func main() {

}
// 1760. 袋子里最少数目的球
// todo
func minimumSize(nums []int, maxOperations int) int {
	f := func(mid int) bool {
		c := 0
		for i := 0; i < len(nums); i++ {
			c += int(math.Ceil(float64(nums[i])/float64(mid)))-1
		}
		return c > maxOperations
	}

	left, right := 1, int(1e9)
	for left < right {
		mid := int(uint(left+right)>>1)
		if f(mid) {
			left = mid+1
		}else {
			right = mid
		}
	}
	return left
}
