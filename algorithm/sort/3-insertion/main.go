package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 生成随机数组
func getRandArray(n int) []int {
	rand.Seed(time.Now().UnixNano())

	var res []int
	for i := 0; i < n; i++ {
		res = append(res, rand.Intn(1e3))
	}
	return res
}

// 插入排序: 每次取一个元素插入到已排序的数据中
func insertSort(nums []int) []int {
	res := nums[0:1]
	for i := 1; i < len(nums); i++ {
		for k := len(res) - 1; k >= 0; k-- {
			// 当插入的元素为最小
			if k == 0 && nums[i] < res[k] {
				res = append([]int{nums[i]}, res...)
				break
			}

			// 当插入的元素为中间, 在k+1点插入元素
			if nums[i] > res[k] {
				// 在slice中插入数据的几种方式
				// 1. 创建临时切片
				// 2. 链式操作 append(res[:k], append([]int{x}, res[k:]...)...)
				// 3. 结合copy操作,先填充元素，再整体移动，后面再插入 res = append(res, 0); copy(res[k+1:], res[k:]); a[k] = val
				res = append(res[:k+1], append([]int{nums[i]}, res[k+1:]...)...)
				break
			}
		}
	}
	return res
}

func main() {
	nums := getRandArray(10)
	fmt.Println(nums)
	fmt.Println(insertSort(nums))
}
