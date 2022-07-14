package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(waysToSplit([]int{1,1,2,2,3,4,5}))
	fmt.Println(waysToSplitV2([]int{1,1,2,2,3,4,5}))
}

// 1712. 将数组分成三个子数组的方案数
// https://leetcode.cn/problems/ways-to-split-array-into-three-subarrays/
// 【前缀和】【二分查找】
// 记录两个分隔的位置为 left, right
// 利用前缀和， S[left] <= S[right] - S[left] <= S[n] - S[right]
// 变形为 2 * S[left] <= S[right]  和  S[left] >= 2*S[right] - S[n]
// 枚举right, 通过二分可以使不等式成立的left的范围，累加这个范围的长度即为答案
// 注意点： 2 <= right < n， 2 * (2*S[right]-S[n]) <= S[right]  =>  3S[right] <= 2S[n]
func waysToSplit(nums []int) int {
	// 构建前缀和
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}
	fmt.Println(nums)

	sum := nums[len(nums)-1]

	ans := 0
	right := 2
	for  right <= len(nums) && 3*nums[right-1] <= 2 * sum {
		l1 := sort.SearchInts(nums[:right-1], 2*nums[right-1]-sum)
		ans += sort.SearchInts(nums[l1:right-1], nums[right-1]/2+1)
		fmt.Println(l1, right, nums[l1:right-1], nums[right-1]/2+1)


		right++
	}

	return ans % (1e9+7)
}


func waysToSplitV2(a []int) (ans int) {
	n := len(a)
	sum := make([]int, n+1)
	for i, v := range a {
		sum[i+1] = sum[i] + v
	}
	fmt.Println(sum)
	for r := 2; r < n && 3*sum[r] <= 2*sum[n]; r++ {
		l1 := sort.SearchInts(sum[1:r], 2*sum[r]-sum[n]) + 1	// 因为前面有一个空零
		ans += sort.SearchInts(sum[l1:r], sum[r]/2+1)
		fmt.Println(l1, r, sum[l1:r], sum[r]/2+1)

	}
	return ans % (1e9 + 7)
}