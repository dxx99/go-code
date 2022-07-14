package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	//fmt.Println(minSumSquareDiff([]int{1,2,3,4}, []int{2,10,20,19}, 0, 0))
	fmt.Println(minSumSquareDiff([]int{1,4,10,12}, []int{5,8,6,9}, 1, 1))
}

// 1838 最高频元素的频次
// https://leetcode.cn/problems/frequency-of-the-most-frequent-element/
func maxFrequency(nums []int, k int) int {
	sort.Ints(nums)

	// 记录频次最高的次数
	max := 1

	total := 0		// 记录当前需要补位的数量
	left, right := 0, 1	// 窗口左右两侧
	for right < len(nums) {
		total += (nums[right] - nums[right-1]) * (right-left)	// 这里比较重要
		for total > k {		// 这里也是关键
			total -= nums[right] - nums[left]
			left++
		}

		if max < right-left+1 {
			max = right-left+1
		}

		right++	// 这里也比较重要
	}

	return max
}

// 2333 最小差值平方和
// https://leetcode.cn/problems/minimum-sum-of-squared-difference/
func minSumSquareDiff(nums1 []int, nums2 []int, k1 int, k2 int) int64 {
	sum := 0
	for i := 0; i < len(nums1); i++ {
		nums1[i] = int(math.Abs(float64(nums1[i]-nums2[i])))
		sum += nums1[i]
	}
	k := k1+k2

	if k >= sum {
		return 0
	}
	
	sort.Slice(nums1, func(i, j int) bool {
		return nums1[i] > nums1[j]
	})
	fmt.Println(nums1)
	f := func(mid int) bool {
		ans := 0
		fmt.Println(mid)
		for i := 0; i < len(nums1); i++ {
			if  nums1[i] > mid {
				ans += nums1[i] - mid
			}
		}
		return ans < k
	}

	nums1 = append(nums1, 0)	// 添加哨兵
	left, right := 0, nums1[0]
	for left < right {
		mid := int(uint(left+right)>>1)
		if f(mid) {
			right = mid
		} else {
			left = mid+1
		}
	}
	fmt.Println(left)

	//求平方和
	var ans int64
	for i := 0; i < len(nums1)-1; i++ {
		if k < nums1[i] - left {
			k -= nums1[i] - left
			ans += int64(left*left)
		}else {
			if k > 0 {
				ans += int64(nums1[i] - k) * int64(nums1[i] - k)
			}else {
				ans += int64(nums1[i]) * int64(nums1[i])
			}
		}
	}


	return ans
}
