package main

import "fmt"

func main() {
	//fmt.Println(minArray([]int{10,1,10,10,10}))
	//fmt.Println(minArray([]int{3,4,5,1,2}))
	//fmt.Println(minArray([]int{2,2,2,0,1}))
	//
	//fmt.Println(minArrayV2([]int{10,1,10,10,10}))
	//fmt.Println(minArrayV2([]int{3,4,5,1,2}))
	fmt.Println(minArrayV2([]int{2,2,2,0,1}))
}

// 二分查找
// 最后的结果都是左右索引都会相等
func minArray(numbers []int) int {
	left, right := 0, len(numbers)-1

	for left < right {
		mid := (right+left)>>1
		if numbers[right] < numbers[mid] {
			left = mid+1
		}else if numbers[right] > numbers[mid] {
			right = mid
		}else {
			right--
		}
	}
	return numbers[left]
}

// 数组【升序--->降序】的二分查找算法
func minArrayV2(numbers []int) int {
	// 正序的数组，直接返回
	if len(numbers) == 1 {
		return numbers[0]
	}

	mid := (len(numbers)-1)>>1

	// 递归使用二分查找
	back := make([]int, 0)
	if numbers[len(numbers)-1] == numbers[mid] {
		back = numbers[:len(numbers)-1]
	} else if  numbers[mid] > numbers[len(numbers)-1] {
		back = numbers[mid+1:]
	}else {
		back = numbers[:mid+1]
	}

	return minArray(back)
}
