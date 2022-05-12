package main

import "fmt"

// p718 最长重复子数组
// 给两个整数数组 nums1 和 nums2 ，返回 两个数组中 公共的 、长度最长的子数组的长度 。
//
// 
//
//示例 1：
//
//输入：nums1 = [1,2,3,2,1], nums2 = [3,2,1,4,7]
//输出：3
//解释：长度最长的公共子数组是 [3,2,1] 。 对应的子数组是连续的
//示例 2：
//
//输入：nums1 = [0,0,0,0,0], nums2 = [0,0,0,0,0]
//输出：5
// 
//
//提示：
//
//1 <= nums1.length, nums2.length <= 1000
//0 <= nums1[i], nums2[i] <= 100
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/maximum-length-of-repeated-subarray
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	fmt.Println(findLength([]int{1,2,3,2,1}, []int{3,2,1,4,7}))
	fmt.Println(findLength([]int{0,0,0,0,0}, []int{0,0,0,0,0}))
}

func findLength(nums1 []int, nums2 []int) int {
	//dp数组定义
	dp := make([][]int, len(nums1)+1)
	for i := range dp {
		dp[i] = make([]int, len(nums2)+1)
	}


	// 初始化
	dp[0][0] = 0
	maxRes := 0
	for i := 1; i <= len(nums1); i++ {
		for j := 1; j <= len(nums2); j++ {
			if nums1[i-1] == nums2[j-1] {	// 这个取值要往后面移一位
				dp[i][j] = dp[i-1][j-1] + 1
			}
			// 保存最大值
			if dp[i][j] > maxRes {
				maxRes = dp[i][j]
			}
		}
	}
	return maxRes
}
