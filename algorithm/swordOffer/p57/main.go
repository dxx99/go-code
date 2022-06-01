package main

import "fmt"

func main() {
	fmt.Println(findContinuousSequence(9))
	fmt.Println(findContinuousSequence(15))
}


func findContinuousSequence(target int) [][]int {
	sum := 0
	ans := make([][]int, 0)
	tmp := make([]int, 0)

	i := 1
	for i < target {
		if sum > target {
			sum -= tmp[0]
			tmp = tmp[1:]
			continue
		}

		if sum == target && len(tmp) >= 2 {
			ans = append(ans, tmp)
		}
		tmp = append(tmp, i)
		sum += i
		i++
	}
	return ans
}
