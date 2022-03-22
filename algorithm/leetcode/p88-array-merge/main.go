package main

import "fmt"

// 88. 合并两个有序数组
// 给你两个按 非递减顺序 排列的整数数组 nums1 和 nums2，另有两个整数 m 和 n ，分别表示 nums1 和 nums2 中的元素数目。
//
//请你 合并 nums2 到 nums1 中，使合并后的数组同样按 非递减顺序 排列。
//
//注意：最终，合并后数组不应由函数返回，而是存储在数组 nums1 中。为了应对这种情况，nums1 的初始长度为 m + n，其中前 m 个元素表示应合并的元素，后 n 个元素为 0 ，应忽略。nums2 的长度为 n 。
//
// 
//
//示例 1：
//
//输入：nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
//输出：[1,2,2,3,5,6]
//解释：需要合并 [1,2,3] 和 [2,5,6] 。
//合并结果是 [1,2,2,3,5,6] ，其中斜体加粗标注的为 nums1 中的元素。
//示例 2：
//
//输入：nums1 = [1], m = 1, nums2 = [], n = 0
//输出：[1]
//解释：需要合并 [1] 和 [] 。
//合并结果是 [1] 。
//示例 3：
//
//输入：nums1 = [0], m = 0, nums2 = [1], n = 1
//输出：[1]
//解释：需要合并的数组是 [] 和 [1] 。
//合并结果是 [1] 。
//注意，因为 m = 0 ，所以 nums1 中没有元素。nums1 中仅存的 0 仅仅是为了确保合并结果可以顺利存放到 nums1 中。
// 
//
//提示：
//
//nums1.length == m + n
//nums2.length == n
//0 <= m, n <= 200
//1 <= m + n <= 200
//-109 <= nums1[i], nums2[j] <= 109
// 
//
//进阶：你可以设计实现一个时间复杂度为 O(m + n) 的算法解决此问题吗？
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/merge-sorted-array
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	merge([]int{1,2,3,0,0,0}, 3, []int{2,5,6}, 3)
	merge([]int{1}, 1, []int{}, 0)
	merge([]int{4,5,6,0,0,0}, 3, []int{1,2,3}, 3)
	fmt.Println("--------------------------------")
	mergeV2([]int{1,2,3,0,0,0}, 3, []int{2,5,6}, 3)
	mergeV2([]int{1}, 1, []int{}, 0)
	mergeV2([]int{4,5,6,0,0,0}, 3, []int{1,2,3}, 3)
}

//TODO: 实现这个方法
func merge(nums1 []int, m int, nums2 []int, n int)  {
	res := make([]int, 0)
	k1, k2 := 0, 0
	for k := 0; k < m+n; k++ {
		if k1 == m {
			res = append(res, nums2[k2:m]...)
			break
		}
		if k2 == n {
			res = append(res, nums1[k1:n]...)
			break
		}
		if nums1[k1] <= nums2[k2] {
			res =append(res, nums1[k1])
			k1 += 1
		}else {
			res = append(res, nums2[k2])
			k2 += 1
		}
	}
	fmt.Println(res)
}

func mergeV2(nums1 []int, m int, nums2 []int, n int)  {
	k1, k2 := m-1, n-1
	for k := m+n-1; k >= 0; k-- {
		if k1 == -1 {
			nums1[k] = nums2[k2]
			k2 -= 1
			continue
		}
		if k2 == -1 {
			break
		}
		if nums1[k1] >= nums2[k2] {
			nums1[k] = nums1[k1]
			k1 -= 1
		}else {
			nums1[k] = nums2[k2]
			k2 -= 1
		}
	}
	fmt.Println(nums1)
}
