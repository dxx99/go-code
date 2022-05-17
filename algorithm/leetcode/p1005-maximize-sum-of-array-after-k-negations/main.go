package main

import (
	"fmt"
	"math"
	"sort"
)

// p1005 k次取反之后最大化的数组和
// 给你一个整数数组 nums 和一个整数 k ，按以下方法修改该数组：
//
//选择某个下标 i 并将 nums[i] 替换为 -nums[i] 。
//重复这个过程恰好 k 次。可以多次选择同一个下标 i 。
//
//以这种方式修改数组后，返回数组 可能的最大和 。
//
// 
//
//示例 1：
//
//输入：nums = [4,2,3], k = 1
//输出：5
//解释：选择下标 1 ，nums 变为 [4,-2,3] 。
//示例 2：
//
//输入：nums = [3,-1,0,2], k = 3
//输出：6
//解释：选择下标 (1, 2, 2) ，nums 变为 [3,1,0,2] 。
//示例 3：
//
//输入：nums = [2,-3,-1,5,-4], k = 2
//输出：13
//解释：选择下标 (1, 4) ，nums 变为 [2,3,-1,5,4] 。
//
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/maximize-sum-of-array-after-k-negations
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	//[-2,5,0,2,-2]
	//3
	fmt.Println(largestSumAfterKNegations([]int{-2,-4,-3}, 4))
	fmt.Println(largestSumAfterKNegations([]int{-2,5,0,2,-2}, 3))
	fmt.Println(largestSumAfterKNegations([]int{-1,-3,-2,-6,-7}, 5))
	fmt.Println(largestSumAfterKNegations([]int{1,3,2,6,7,9}, 3))
	fmt.Println(largestSumAfterKNegations([]int{4,2,3}, 1))
	fmt.Println(largestSumAfterKNegations([]int{3,-1,0,2}, 3))
	fmt.Println(largestSumAfterKNegations([]int{2,-3,-1,5,-4}, 2))
}

// 贪心算法，将指针指向绝对值最小的元素
func largestSumAfterKNegations(nums []int, k int) int {
	sort.Ints(nums)

	i := 0
	for k > 0 {
		// 已经到了数组最后面
		if i == len(nums) {
			if k % 2 == 1 {
				nums[i-1] = -nums[i-1]
			}
			break
		}

		// 前面的元素开始为非负数
		if nums[i] >= 0 {
			if k % 2 == 1 {
				if i > 0 && int(math.Abs(float64(nums[i]))) > int(math.Abs(float64(nums[i-1]))) {
					nums[i-1] = -nums[i-1]
				}else {
					nums[i] = -nums[i]
				}
			}
			break
		}

		// 如果元素为负数
		if nums[i] < 0 {
			nums[i] = -nums[i]
			i++
			k--
		}

	}

	ans := 0
	for _, num := range nums {
		ans += num
	}
	return ans
}
