package main

import (
	"fmt"
	"sort"
)

func main() {
	//arr := [][]int{
	//{72,98},{62,27},{32,7},{71,4},{25,19},{91,30},{52,73},{10,9},{99,71},
	//{47,22},{19,30},{80,63},{18,15},{48,17},{77,16},{46,27},{66,87},{55,84},
	//{65,38},{30,9},{50,42},{100,60},{75,73},{98,53},{22,80},{41,61},{37,47},
	//{95,8},{51,81},{78,79},{57,95},
	//}
	//fmt.Println(minimumLines(arr))
	////
	fmt.Println(minimumLines([][]int{{1,7},{2,6},{3,5},{4,4},{5,4},{6,3},{7,2},{8,1}}))
	fmt.Println(minimumLines([][]int{{3,4},{1,2},{7,8},{2,3}}))
	fmt.Println(minimumLines([][]int{{2,1},{1,2}}))
	fmt.Println(minimumLines([][]int{{1,1},{500000000,499999999},{1000000000,999999998}}))
}

// 注意斜率的计算公式，吐了
func minimumLines(stockPrices [][]int) int {
	if len(stockPrices) <= 1 {
		return 0
	}
	sort.Slice(stockPrices, func(i, j int) bool {
		return stockPrices[i][0] < stockPrices[j][0]
	})

	ans := 1
	for i := 2; i < len(stockPrices); i++ {
		x1, y1 := stockPrices[i][0] - stockPrices[i-1][0], stockPrices[i][1] - stockPrices[i-1][1]
		x2, y2 := stockPrices[i-1][0] - stockPrices[i-2][0], stockPrices[i-1][1] - stockPrices[i-2][1]
		if x1*y2 != y1*x2 {	// 斜率计算，有精度问题
			ans++
		}
	}
	return ans
}
