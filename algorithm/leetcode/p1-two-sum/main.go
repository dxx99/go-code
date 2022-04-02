package main

import (
	"fmt"
	"sort"
)

// p1 两数之和
// 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
//
//你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
//
//你可以按任意顺序返回答案。
//
// 
//
//示例 1：
//
//输入：nums = [2,7,11,15], target = 9
//输出：[0,1]
//解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。
//示例 2：
//
//输入：nums = [3,2,4], target = 6
//输出：[1,2]
//示例 3：
//
//输入：nums = [3,3], target = 6
//输出：[0,1]
// 
//
//提示：
//
//2 <= nums.length <= 104
//-109 <= nums[i] <= 109
//-109 <= target <= 109
//只会存在一个有效答案
//进阶：你可以想出一个时间复杂度小于 O(n2) 的算法吗？
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/two-sum
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	fmt.Println(twoSum([]int{2,7,11,15}, 9))
	fmt.Println(twoSum([]int{3,2,4}, 6))
	fmt.Println(twoSum([]int{3,3}, 6))
	fmt.Println("V2.........................")

	fmt.Println(twoSumV2([]int{2,7,11,15}, 9))
	fmt.Println(twoSumV2([]int{3,2,4}, 6))
	fmt.Println(twoSumV2([]int{3,3}, 6))
}

func twoSum(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{0, 0}
	}

	// map存储
	m := make(map[int][]int)
	for k, v := range nums {
		if s, ok := m[v]; ok {
			m[v] = append(s, k)
		} else {
			m[v] = []int{k}
		}

	}

	// 排序
	sort.Sort(sort.IntSlice(nums))
	left, right := 0, len(nums)-1

	for left < right {
		if nums[left] + nums[right] > target {
			right--
		}else if nums[left] + nums[right] < target {
			left++
		}else {
			if nums[left] == nums[right]  {
				return m[nums[left]]
			}
			return []int{m[nums[left]][0], m[nums[right]][0]}
		}
	}

	return []int{0, 0}
}

func twoSumV2(nums []int, target int) []int {
	m := make(map[int]int)
	for k, v := range nums {
		if p, ok := m[target - v]; ok {
			return []int{p, k}
		}
		m[v] = k
	}
	return []int{0, 0}
}
