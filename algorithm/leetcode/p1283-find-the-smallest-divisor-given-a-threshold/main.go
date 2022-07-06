package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(smallestDivisor([]int{44,22,33,11,1}, 5))
	fmt.Println(smallestDivisor([]int{1,2,5,9}, 6))
	fmt.Println(smallestDivisor([]int{2,3,5,7,11}, 11))
	fmt.Println(smallestDivisor([]int{19}, 5))
}

func smallestDivisor(nums []int, threshold int) int {
	left, right := 1, int(1e6)

	f := func(mid int) bool {
		ans := 0
		for i := 0; i < len(nums); i++ {
			ans += int(math.Ceil(float64(nums[i])/float64(mid)))
		}
		return ans > threshold
	}

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
