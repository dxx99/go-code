package main

import "fmt"

func main() {
	//fmt.Println(minSubArrayLen(7, []int{2,3,1,2,4,3}))
	//fmt.Println(minSubArrayLen(11, []int{1,1,1,1,1,1,1,1}))
	fmt.Println(minSubArrayLen(15, []int{1,2,3,4,5}))
}

func minSubArrayLen(target int, nums []int) int {
	ans := len(nums)+1
	curSum := nums[0]
	left, right := 0, 1
	for  {
		// 终止条件[到达了最右端，且没有合大于目标数了]
		if right == len(nums) && curSum < target {
			break
		}

		if curSum < target && right < len(nums) {
			curSum += nums[right]
			right++
		}

		// 左节点右移
		if curSum >= target {
			// 记录元素个数
			if right - left < ans {
				ans = right - left
			}
			curSum -= nums[left]
			left++
		}
	}
	if ans == len(nums)+1 || ans <= 0 {
		return 0
	}
	return ans
}
