package main

import "fmt"

// p46 全排列
// 给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
//
// 
//
//示例 1：
//
//输入：nums = [1,2,3]
//输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
//示例 2：
//
//输入：nums = [0,1]
//输出：[[0,1],[1,0]]
//示例 3：
//
//输入：nums = [1]
//输出：[[1]]
// 
//
//提示：
//
//1 <= nums.length <= 6
//-10 <= nums[i] <= 10
//nums 中的所有整数 互不相同
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/permutations
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	fmt.Println(permute([]int{1,2,3}))
	fmt.Println(permute([]int{0,1}))
	fmt.Println(permute([]int{1}))
}

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	track := make([]int, 0)
	used := [21]int{}  	// 使用数组来存储用过的值

	var backtracking func()
	backtracking = func() {
		// 终止条件
		if len(track) == len(nums) {
			tmp := make([]int, len(track))
			copy(tmp, track)
			res = append(res, tmp)
			return
		}

		//遍历
		for i := 0; i < len(nums); i++ {
			// 在竖列上这个元素已经被使用了
			if used[nums[i]+10] != 0 {
				continue
			}

			used[nums[i]+10]++
			track = append(track, nums[i])

			//递归
			backtracking()

			// 回溯
			used[nums[i]+10]--
			track = track[:len(track)-1]
		}
	}

	backtracking()
	return res
}
