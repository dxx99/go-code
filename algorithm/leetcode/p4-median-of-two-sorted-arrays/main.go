package main

import (
	"fmt"
)

// p4 寻找两个正序数组的中位数
// 给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。
//
//算法的时间复杂度应该为 O(log (m+n)) 。
//
// 
//
//示例 1：
//
//输入：nums1 = [1,3], nums2 = [2]
//输出：2.00000
//解释：合并数组 = [1,2,3] ，中位数 2
//示例 2：
//
//输入：nums1 = [1,2], nums2 = [3,4]
//输出：2.50000
//解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5
// 
//
// 
//
//提示：
//
//nums1.length == m
//nums2.length == n
//0 <= m <= 1000
//0 <= n <= 1000
//1 <= m + n <= 2000
//-106 <= nums1[i], nums2[i] <= 106
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/median-of-two-sorted-arrays
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	fmt.Println(findMedianSortedArrays([]int{1,3}, []int{2}))
	fmt.Println(findMedianSortedArrays([]int{1,2}, []int{3,4}))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1, l2 := len(nums1), len(nums2)
	nl := l1+l2
	newArr := make([]int, 0)

	// 合并数组
	k1, k2 := 0, 0
	for i := 0; i < nl; i++ {
		if k1 == l1 {
			newArr = append(newArr, nums2[k2:]...)
			break
		}
		if k2 == l2 {
			newArr = append(newArr, nums1[k1:]...)
			break
		}

		if nums2[k2] > nums1[k1] {
			newArr = append(newArr, nums1[k1])
			k1++
		}else{
			newArr = append(newArr, nums2[k2])
			k2++
		}
	}
	fmt.Println(newArr)

	// 返回中位数
	if (nl) % 2 == 0 {
		return (float64(newArr[nl/2-1]) + float64(newArr[nl/2]))/2
	}else {
		return float64(newArr[(nl-1)/2])
	}
}

