package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	// ["64333639502","65953866768","17845691654","87148775908","58954177897","70439926174","48059986638","47548857440","18418180516","06364956881","01866627626","36824890579","14672385151","71207752868"]
	//[[9,4],[6,1],[3,8],[12,9],[11,4],[4,9],[2,7],[10,3],[13,1],[13,1],[6,1],[5,10]]
	//nums := []string{"64333639502","65953866768","17845691654","87148775908","58954177897","70439926174","48059986638","47548857440","18418180516","06364956881","01866627626","36824890579","14672385151","71207752868"}
	//queries := [][]int{{9,4},{6,1},{3,8},{12,9},{11,4},{4,9},{2,7},{10,3},{13,1},{13,1},{6,1},{5,10}}
	//
	//fmt.Println(smallestTrimmedNumbers(nums, queries))
	fmt.Println(smallestTrimmedNumbers([]string{"24","37","96","04"}, [][]int{{2,1},{2,2}}))
	//fmt.Println(smallestTrimmedNumbers([]string{"102","473","251","814"}, [][]int{{1,1},{2,3},{4,2},{1,2}}))
	//fmt.Println(maximumSum([]int{18,43,36,13,7}))
	//fmt.Println(maximumSum([]int{10,12,19,14}))
}

//1.
func numberOfPairs(nums []int) []int {
	sort.Ints(nums)
	ans := 0
	left := -1
	for i := 0; i < len(nums); i++ {
		if left != -1 && nums[i] == left {
			left = -1
			ans++
		}else {
			left = nums[i]
		}
	}
	return []int{ans, len(nums)-2*ans}
}

//2.
func maximumSum(nums []int) int {
	bs := func(n int) int {
		ans := 0
		for n > 9 {
			ans += n % 10
			n = n / 10
		}
		ans += n
		return ans
	}

	sort.Ints(nums)
	bitSum := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		bitSum[i] = bs(nums[i])
	}
	fmt.Println(nums, "\n", bitSum)
	res := -1
	for i := len(nums)-1; i > 0 ; i-- {
		for j := i-1; j >= 0; j-- {
			if bitSum[i] == bitSum[j] {
				if nums[i] + nums[j] > res {
					res = nums[i] + nums[j]
				}
			}
		}
	}
	return res
}

//3.
func smallestTrimmedNumbers(nums []string, queries [][]int) []int {
	hash := make(map[int][][]int, len(nums[0]))
	for i := range hash {
		hash[i] = make([][]int, 0)
	}
	nl := len(nums[0])
	for i := 0; i < len(nums); i++ {
		for j := 1; j <= nl ; j++ {
			tmp, _ := strconv.Atoi(nums[i][nl-j:])
			ns, _ := strconv.Atoi(nums[i])
			hash[j] = append(hash[j], []int{tmp, i, ns})
		}
	}

	for k := range hash {
		item := hash[k]
		sort.Slice(item, func(i, j int) bool {
			if item[i][0] == item[j][0] {
				return item[i][2] < item[j][2]
			}
			return item[i][0] < item[j][0]
		})
		hash[k] = item
	}

	ans := make([]int, len(queries))
	for i, query := range queries {
		if i == 1 {
			fmt.Println(hash[i])
			fmt.Println(query, query[0]-1)
		}
		ans[i] = hash[query[1]][query[0]-1][1]
	}
	return ans
}

//4.

