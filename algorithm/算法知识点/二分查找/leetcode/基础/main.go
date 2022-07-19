package main

import (
	"math"
	"math/rand"
	"sort"
	"time"
)

//162
//154
//528
//1508
//1574
//1292
//1498
//981
//1300
//1802
//1901
//1146
//1488
func main() {
	obj := Constructor([]int{1})
	obj.PickIndex()
	obj.PickIndex()
}

// 162 寻找峰值
// https://leetcode.cn/problems/find-peak-element/
// 思路：
//	1. 寻找最大值, 遍历取最大值即可
//  2. 根据 nums[i-1], nums[i], nums[i+1] 三者的关系决定向哪个方向走
//		- 如果nums[i]大于左右两边的值，直接返回
//		- 如果上坡, 也就是 nums[i-1]<nums[i]<nums[i+1], 最大值在右边，left=mid+1
//		- 如果下坡，也就是 nums[i-1]>nums[i]>nums[i+1], 最大值在左边，right=mid-1
//		- 如果低谷，也就是 nums[i-1]>nums[i]<nums[i+1], 也就是两侧都可以走
//	- 总结：
//		- nums[i] < nums[i+1], 右边
//		- nums[i] > nums[i+1], 左边
func findPeakElement(nums []int) int {
	get := func(i int) int {
		if i == -1 || i == len(nums) {
			return math.MinInt
		}
		return nums[i]
	}

	left, right := 0, len(nums)-1
	for{
		mid := int(uint(left+right)>>1)
		if get(mid-1) < get(mid) && get(mid) > get(mid+1) {
			return mid
		}else if get(mid) < get(mid+1) {
			left = mid+1
		}else {
			right = mid-1
		}
	}
}

// 153 寻找旋转排序数组中的最小值
func findMinI(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := int(uint(left+right)>>1)
		if nums[mid] > nums[right] {	// 思考，与右边元素比较，判断最小值在那个区间
			left = mid+1
		}else {
			right = mid
		}
	}
	return nums[left]
}

// 154 寻找旋转排序数组中的最小值II
// 思路：
//	1. 枚举mid未知的值与right节点的所有可能情况
//		- nums[mid] > nums[right]  最小值一定在mid的右边
//		- nums[mid] = nums[right] 忽略二分查找，右节点向左移一位
//		- nums[mid] < nums[right] 最小值一定在mid的左边
func findMinII(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := int(uint(left+right)>>1)
		if nums[mid] > nums[right] {
			left = mid+1
		} else if nums[mid] < nums[right] {
			right = mid	// 只能移动到这个位置
		}else {		// 相等要额外处理，这里就不能用二分了，只能移动
			right--
		}
	}
	return nums[left]
}


// Solution 前缀和存储
type Solution struct {
	sum []int
}
func Constructor(w []int) Solution {
	sum := make([]int, len(w))
	sum[0] = w[0]
	for i := 1; i < len(sum); i++ {
		sum[i] += sum[i-1]+w[i]
	}
	return Solution{
		sum: sum,
	}
}

// PickIndex
// 528 按权重随机选择
func (s *Solution) PickIndex() int {
	rand.Seed(time.Now().UnixMicro())
	x := sort.SearchInts(s.sum, rand.Intn(s.sum[len(s.sum)-1]))
	return x
}

// 1508 子数组和排序后的区间
func rangeSum(nums []int, n int, left int, right int) int {
	sums := make([]int, n*(n+1)/2)
	index := 0
	for i := 0; i < n; i++ {
		total := 0
		for j := i; j < n; j++ {
			total += nums[j]
			sums[index] = total
			index++
		}
	}

	sort.Ints(sums)
	ans := 0
	for i := left-1; i < right; i++ {
		ans = (ans+sums[i]) % (1e9+7)
	}

	return ans
}




// 1901 矩阵中寻找峰值
// https://leetcode.cn/problems/find-a-peak-element-ii/
// 思路：
//	1. 左右找最大的
//	2. 上下找峰值
func findPeakGrid(mat [][]int) []int {
	getMaxIndex := func(arr []int) int {
		m, index := arr[0], 0
		for i := 1; i < len(arr); i++ {
			if arr[i] > m {
				m = arr[i]
				index = 0
			}
		}
		return index
	}

	m := len(mat)
	left, right := 0, m-1
	for left < right {
		mid := int(uint(left+right)>>1)
		// 左右判断
		maxIndex := getMaxIndex(mat[mid])

		// 上下判断
		if mat[mid][maxIndex] > mat[mid+1][maxIndex]  {
			right = mid
		}else {	// 相等情况要向下移
			left = mid+1
		}
	}

	maxIndex := getMaxIndex(mat[left])
	return []int{left, maxIndex}
}






