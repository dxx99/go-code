package main

import "fmt"

func main() {
	fmt.Println(mySqrt(4))
}

func mySqrt(x int) int {
	left, right := 1, x
	for left <= right {
		mid := (left+right)>>1
		if mid*mid == x {
			return mid
		} else if mid*mid > x {
			right = mid -1
		} else {
			left = mid + 1
		}
	}
	return left
}
