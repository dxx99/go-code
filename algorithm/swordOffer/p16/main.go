package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Pow(2.1, 3))
	fmt.Println(myPow(2.1, 3))
	fmt.Println(math.Pow(0.00001, 2147483647))
}

// 递归函数，分治求解
func myPow(x float64, n int) float64 {
	var quickMul func(num float64, n int) float64
	quickMul = func(num float64, n int) float64 {
		if n == 0 {
			return 1
		}

		y := quickMul(num, n/2)
		if n % 2 == 0 {
			return y * y
		} else {
			return y * y * x
		}
	}

	if n >= 0 {
		return quickMul(x, n)
	}
	return 1.0/quickMul(x, -n)
}
