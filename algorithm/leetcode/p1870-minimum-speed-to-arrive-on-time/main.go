package main

import "math"

func main() {

}

func minSpeedOnTime(dist []int, hour float64) int {
	f := func(mid int) bool {
		used := 0.0
		for i := 0; i < len(dist)-1; i++ {
			used += math.Ceil(float64(dist[i])/float64(mid))
		}
		used += float64(dist[len(dist)-1])/float64(mid)

		return used > hour
	}

	left, right := 1, int(1e5)
	for left < right {
		mid := int(uint(left+right)>>1)
		if f(mid) {
			left = mid+1
		}else {
			right = mid
		}
	}

	// 临界值判断
	if left == int(1e5) && f(left) {
		return -1
	}
	return left
}
