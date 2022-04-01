package main

import (
	"fmt"
	"sort"
)

// 278. 第一个错误的版本
// 你是产品经理，目前正在带领一个团队开发新的产品。不幸的是，你的产品的最新版本没有通过质量检测。由于每个版本都是基于之前的版本开发的，所以错误的版本之后的所有版本都是错的。
//
//假设你有 n 个版本 [1, 2, ..., n]，你想找出导致之后所有版本出错的第一个错误的版本。
//
//你可以通过调用 bool isBadVersion(version) 接口来判断版本号 version 是否在单元测试中出错。实现一个函数来查找第一个错误的版本。你应该尽量减少对调用 API 的次数。
//
// 
//示例 1：
//
//输入：n = 5, bad = 4
//输出：4
//解释：
//调用 isBadVersion(3) -> false
//调用 isBadVersion(5) -> true
//调用 isBadVersion(4) -> true
//所以，4 是第一个错误的版本。
//示例 2：
//
//输入：n = 1, bad = 1
//输出：1
// 
//
//提示：
//
//1 <= bad <= n <= 231 - 1
//
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/first-bad-version
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	fmt.Println(firstBadVersion(21))
	fmt.Println(firstBadVersionV2(21))
	fmt.Println(firstBadVersionV3(21))
}

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */
func isBadVersion(version int) bool {
	bad := 11
	if version >= bad {
		return true
	}
	return false
}

// 变种二分查找
func firstBadVersion(n int) int {
	mid := (1 + n) / 2
	min, max := 1, n
	for i := 0; i < n; i++ {

		if !isBadVersion(mid) {		// false区间, 表示全部都是正常的版本
			if isBadVersion(mid + 1) {		// false-true 🔗表示第一个坏版本
				return mid+1
			}
			min = mid + 1
		} else {     				// true区间，表示全部都是错误的版本
			if !isBadVersion(mid - 1) {		// false-true 🔗表示第一个坏版本
				return mid
			}
			max = mid - 1
		}
		mid = (min + max) / 2

	}
	return 0
}

// 优化代码，使用双指针处理
func firstBadVersionV2(n int) int {
	left, right := 1, n
	for left < right {	// 直到左右两边相关才退出
		mid := (left + right) / 2

		if !isBadVersion(mid) {		// false, 肯定在右边+1
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

// 直接调用sort包的Search函数
func firstBadVersionV3(n int) int {
	b := sort.Search(n, func(version int) bool {
		return isBadVersion(version)
	})

	// 排除掉没有坏版本的结果
	if b == n && !isBadVersion(b) {
		return 0
	}
	return b
}