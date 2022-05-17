package main

import "fmt"

// p376 摆动序列
// 如果连续数字之间的差严格地在正数和负数之间交替，则数字序列称为 摆动序列 。第一个差（如果存在的话）可能是正数或负数。仅有一个元素或者含两个不等元素的序列也视作摆动序列。
//
//例如， [1, 7, 4, 9, 2, 5] 是一个 摆动序列 ，因为差值 (6, -3, 5, -7, 3) 是正负交替出现的。
//
//相反，[1, 4, 7, 2, 5] 和 [1, 7, 4, 5, 5] 不是摆动序列，第一个序列是因为它的前两个差值都是正数，第二个序列是因为它的最后一个差值为零。
//子序列 可以通过从原始序列中删除一些（也可以不删除）元素来获得，剩下的元素保持其原始顺序。
//
//给你一个整数数组 nums ，返回 nums 中作为 摆动序列 的 最长子序列的长度 。
//
// 
//
//示例 1：
//
//输入：nums = [1,7,4,9,2,5]
//输出：6
//解释：整个序列均为摆动序列，各元素之间的差值为 (6, -3, 5, -7, 3) 。
//示例 2：
//
//输入：nums = [1,17,5,10,13,15,10,5,16,8]
//输出：7
//解释：这个序列包含几个长度为 7 摆动序列。
//其中一个是 [1, 17, 10, 13, 10, 16, 8] ，各元素之间的差值为 (16, -7, 3, -3, 6, -8) 。
//示例 3：
//
//输入：nums = [1,2,3,4,5,6,7,8,9]
//输出：2
// 
//
//提示：
//
//1 <= nums.length <= 1000
//0 <= nums[i] <= 1000
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/wiggle-subsequence
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	fmt.Println(wiggleMaxLength([]int{1,7,4,9,2,5}))
	fmt.Println(wiggleMaxLength([]int{1,17,5,10,13,15,10,5,16,8}))
	fmt.Println(wiggleMaxLength([]int{1,2,3,4,5,6,7,8,9}))
	fmt.Println(wiggleMaxLength([]int{0,0}))
	fmt.Println(wiggleMaxLength([]int{84}))

	fmt.Println("----------------")
	fmt.Println(wiggleMaxLengthV2([]int{1,7,4,9,2,5}))
	fmt.Println(wiggleMaxLengthV2([]int{1,17,5,10,13,15,10,5,16,8}))
	fmt.Println(wiggleMaxLengthV2([]int{1,2,3,4,5,6,7,8,9}))
	fmt.Println(wiggleMaxLengthV2([]int{0,0}))
	fmt.Println(wiggleMaxLengthV2([]int{84}))
}

// 只有发生拐点的时候才需要记录，然后中间的元素都不能算
func wiggleMaxLength(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	wNum := 1

	// 如果操作为1, 下一个操作表示增加， 如果为2表示操作减少
	op := 0
	for i := 1; i < len(nums); i++ {
		if (op == 1 || op == 0) && nums[i]  > nums[i-1] {
			op = 2
			wNum++
		}

		if (op == 2 || op == 0) && nums[i] < nums[i-1] {
			op = 1
			wNum++
		}
	}

	return wNum
}


// 动态规划求解
func wiggleMaxLengthV2(nums []int) int {

	// dp[i][0] 表示第i个元素作为山峰，摆动子序列的最大长度
	// dp[i][1] 表示第i个元素作为山谷，摆动子序列的最大长度
	dp := make([][]int, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, 2)
	}

	// 初始化
	dp[0][0], dp[0][1] = 1, 1

	// 遍历
	for i := 1; i < len(nums); i++ {
		dp[i][0], dp[i][1] = 1, 1
		for j := 0; j < i; j++ {
			if nums[j] > nums[i] {	// 波谷
				dp[i][1] = max(dp[i][1], dp[j][0] + 1)
			}
			if nums[j] < nums[i] {
				dp[i][0] = max(dp[i][0], dp[j][1] + 1)
			}
		}
	}

	return max(dp[len(nums)-1][0], dp[len(nums)-1][1])
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
