package main

import (
	"fmt"
	"sort"
)

// p455 分发饼干
// 假设你是一位很棒的家长，想要给你的孩子们一些小饼干。但是，每个孩子最多只能给一块饼干。
//
//对每个孩子 i，都有一个胃口值 g[i]，这是能让孩子们满足胃口的饼干的最小尺寸；并且每块饼干 j，都有一个尺寸 s[j] 。如果 s[j] >= g[i]，我们可以将这个饼干 j 分配给孩子 i ，这个孩子会得到满足。你的目标是尽可能满足越多数量的孩子，并输出这个最大数值。
//
// 
//示例 1:
//
//输入: g = [1,2,3], s = [1,1]
//输出: 1
//解释:
//你有三个孩子和两块小饼干，3个孩子的胃口值分别是：1,2,3。
//虽然你有两块小饼干，由于他们的尺寸都是1，你只能让胃口值是1的孩子满足。
//所以你应该输出1。
//示例 2:
//
//输入: g = [1,2], s = [1,2,3]
//输出: 2
//解释:
//你有两个孩子和三块小饼干，2个孩子的胃口值分别是1,2。
//你拥有的饼干数量和尺寸都足以让所有孩子满足。
//所以你应该输出2.
// 
//
//提示：
//
//1 <= g.length <= 3 * 104
//0 <= s.length <= 3 * 104
//1 <= g[i], s[j] <= 231 - 1
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/assign-cookies
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	fmt.Println(findContentChildren([]int{1,2,3}, []int{1,1}))
	fmt.Println(findContentChildren([]int{1,2}, []int{1,2,3}))

	fmt.Println(findContentChildrenV2([]int{1,2,3}, []int{1,1}))
	fmt.Println(findContentChildrenV2([]int{1,2}, []int{1,2,3}))

}

func findContentChildren(g []int, s []int) int {
	num := 0
	sort.Ints(s)

	// 每个孩子的胃口值
	for _, item := range g {
		if len(s) == 0 {
			return num
		}
		for i := 0; i < len(s); i++ {
			if s[i] >= item  {
				num++
				s = append(s[:i], s[i+1:]...)
				break
			}
		}
	}
	return num
}

// 优化贪心算法
// 思路1：优先考虑饼干，小饼干应该给小胃口， 饼干从前往后遍历
// 思路2：优先考虑胃口，大饼干优先满足胃口大的人员，胃口从后往前遍历
func findContentChildrenV2(g []int, s []int) int {
	num := 0

	sort.Ints(g)
	sort.Ints(s)

	right := len(s)-1
	for i := len(g)-1; i >= 0 ; i-- {
		if right >= 0 && s[right] >= g[i] {		//饼干大于胃口
			num++
			right--
		}
	}
	return num
}
