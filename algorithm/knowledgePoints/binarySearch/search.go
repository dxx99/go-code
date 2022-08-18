package main

import "fmt"

// binarySearch
// 左右边界的二分查找常用的左闭右开的方法来查找
// 二分思维的精髓就是：通过已知信息尽可能多地收缩（折半）搜索空间
func main() {
	fmt.Println("binarySearch: ", binarySearch([]int{1,2,3,4,4,4,4,5,6}, 4))
	fmt.Println("right: ", rightBound([]int{1,2,3,4,4,4,5,6}, 4))
	fmt.Println("rightV2: ", rightBoundV2([]int{1,2,3,4,4,4,5,6}, 4))
	fmt.Println("left: ", leftBound([]int{1,2,3,4,4,4,5,6}, 4))
	fmt.Println("leftV2: ", leftBoundV2([]int{1,2,3,4,4,4,5,6}, 4))
}

// 左闭右闭
func binarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := int(uint(left+right)>>1)
		if nums[mid] < target {
			left = mid+1
		}else if nums[mid] > target {
			right = mid-1
		}else {
			return mid	//找到，直接退出
		}
	}
	return -1	//没有找到
}


// 查找最左边的值，也就是第一个等于目标值的索引
// [left, right) 左闭右开版本
func leftBound(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := int(uint(left+right)>>1)
		if nums[mid] < target {		// 相等之后就不会右移，因此找到的是左边界
			left = mid+1
		}else {
			right = mid
		}
	}

	// 找不到返回-1，优化返回结果集
	if left == len(nums) || nums[left] != target {
		return -1
	}
	return left
}
// [left, right] 左闭右闭版本
func leftBoundV2(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := int(uint(left+right)>>1)
		if nums[mid] < target {
			left = mid+1			// [mid+1, right]
		}else if nums[mid] > target {
			right = mid-1			// [left, mid-1]
		}else if nums[mid] == target {
			right = mid-1			// 【注意】: 收缩右边界，这样才能退出
		}
	}

	// 找不到返回-1，优化返回结果集
	if left >= len(nums) || nums[left] != target {
		return -1
	}
	return left
}



// 查找最右边的值，也就是最后一个等于目标值的索引
// [left, right) 左闭右开版本
func rightBound(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := int(uint(left+right)>>1)
		if nums[mid] <= target {	// 注意这里的等号，也需要右移
			left = mid+1
		}else {
			right = mid
		}
	}

	// 找不到返回-1，优化返回结果集
	if left <= 0 || nums[left-1] != target {
		return -1
	}
	return left-1
}
// [left, right] 左闭右闭版本
func rightBoundV2(nums []int, target int) int {
	left, right := 0, len(nums)
	for left <= right {
		mid := int(uint(left+right)>>1)
		if nums[mid] < target {
			left = mid+1
		}else if nums[mid] > target {
			right = mid-1
		}else if nums[mid] == target {
			left = mid+1		// 【注意】: 收缩左边界，这样才能退出
		}
	}

	// 找不到返回-1，优化返回结果集
	if left <= 0 || nums[left-1] != target {
		return -1
	}
	return left-1
}

