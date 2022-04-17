package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(getUsedTime(4))
	fmt.Println(minimumRounds([]int{2,2,3,3,2,4,4,4,4,4}))
	fmt.Println(minimumRounds([]int{9,9,11,9,11,22,22,22,22,55,55}))
}

func minimumRounds(tasks []int) int {
	hashMap := make(map[int]int, 0)
	for _, num := range tasks {
		hashMap[num]++
	}
	usedTime := 0
	sort.Ints(tasks)
	stack := make([]int, 0)
	for _, e := range tasks {
		if num, ok := hashMap[e]; ok{

			if num == 1 {
				// 出栈操作
				eNum := 1
				for i := len(stack)-1; i >=0 ; i-- {
					if stack[i] == e {
						eNum++
					}else{
						stack = stack[:i+1]
						break
					}
				}
				if eNum == 1  {
					return -1
				}
				//判断使用多少次
				usedTime += getUsedTime(eNum)
			}else {
				// 入栈操作
				hashMap[e]--
				stack = append(stack, e)
			}
		}
	}

	return usedTime
}

func getUsedTime(num int) int {
	if num == 1 {
		return -1
	}
	if num == 2 || num == 3 {
		return 1
	}

	t := num / 3
	if num % 3 != 0 {
		t++
	}
	return t
}
