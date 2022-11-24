package main

import (
	"fmt"
	"sort"
)

// p870. 优势洗牌
// 给定两个大小相等的数组 nums1 和 nums2，nums1 相对于 nums2 的优势可以用满足 nums1[i] > nums2[i] 的索引 i 的数目来描述。
//
// 返回 nums1 的任意排列，使其相对于 nums2 的优势最大化。
//
// 示例 1：
//
// 输入：nums1 = [2,7,11,15], nums2 = [1,10,4,11]
// 输出：[2,11,7,15]
// 示例 2：
//
// 输入：nums1 = [12,24,8,32], nums2 = [13,25,32,11]
// 输出：[24,32,8,12]
//
// 提示：
//
// 1 <= nums1.length <= 105
// nums2.length == nums1.length
// 0 <= nums1[i], nums2[i] <= 109
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode.cn/problems/advantage-shuffle
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	fmt.Println(advantageCount([]int{12, 24, 8, 32}, []int{13, 25, 32, 11}))
	fmt.Println(advantageCountV2([]int{12, 24, 8, 32}, []int{13, 25, 32, 11}))
}

// 思路：
//  1. 先排序
//  2. 查找, 刚好大的，没有就用最小的
//  3. 组合
func advantageCount(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	res := make([]int, 0)
	for _, num := range nums2 {
		k := sort.SearchInts(nums1, num+1)
		if k >= len(nums1) {
			k = 0
		}
		res = append(res, nums1[k])
		nums1 = append(nums1[:k], nums1[k+1:]...)
	}
	return res
}

// todo: 贪心算法
func advantageCountV2(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)

	// 排序nums2, 按从大到小的顺序排序
	idx := make([]int, len(nums2))
	for i := 0; i < len(idx); i++ {
		idx[i] = i
	}
	sort.Slice(idx, func(i, j int) bool {
		return nums2[idx[i]] < nums2[idx[j]]
	})

	left, right := 0, len(nums1)-1
	ans := make([]int, len(nums1))
	for _, num := range nums1 {
		if num > nums2[idx[left]] {
			ans[idx[left]] = num
			left++
		} else {
			ans[idx[right]] = num
			right--
		}
	}
	return ans
}
