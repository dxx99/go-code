package main

import "fmt"

// p763 划分字母区间
// 字符串 S 由小写字母组成。我们要把这个字符串划分为尽可能多的片段，同一字母最多出现在一个片段中。返回一个表示每个字符串片段的长度的列表。
//
// 
//
//示例：
//
//输入：S = "ababcbacadefegdehijhklij"
//输出：[9,7,8]
//解释：
//划分结果为 "ababcbaca", "defegde", "hijhklij"。
//每个字母最多出现在一个片段中。
//像 "ababcbacadefegde", "hijhklij" 的划分是错误的，因为划分的片段数较少。
// 
//
//提示：
//
//S的长度在[1, 500]之间。
//S只包含小写字母 'a' 到 'z' 。
//
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/partition-labels
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	fmt.Println(partitionLabels("ababcbacadefegdehijhklij"))
	fmt.Println(partitionLabelsV2("ababcbacadefegdehijhklij"))
}

// 贪心算法 + 优化代码
func partitionLabelsV2(s string) []int {
	ans := make([]int, 0)

	// 存储元素最后一次出现的位置
	hash := make([]int, 26)
	for i := 0; i < len(s); i++ {
		hash[s[i]-'a'] = i
	}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		k := s[i] - 'a'

		// 得到当前区间最远的距离
		if hash[k] > end {
			end = hash[k]
		}

		if i == end {
			ans = append(ans, end-start+1)
			start = end+1
		}
	}
	return ans
}


// 贪心算法
func partitionLabels(s string) []int {
	ans := make([]int, 0)

	// 使用数组hash来存储结果
	hash := make([]int, 26)
	for i := 0; i < len(s); i++ {
		hash[s[i]-'a']++
	}

	// 用来判断数组是否存在元素
	isExistElement := func(arr []bool) bool {
		for i := 0; i < len(arr); i++ {
			if arr[i] == true {
				return true
			}
		}
		return false
	}

	cur := 0
	tmp := make([]bool, 26)
	for i := 0; i < len(s); i++ {
		cur++
		k := s[i]-'a'
		hash[k]--

		if hash[k] > 0 {
			tmp[k] = true
		} else {
			tmp[k] = false
		}

		// 如果临时数组中不存在元素， 尽可能分隔数据
		if !isExistElement(tmp) {
			ans = append(ans, cur)
			tmp = make([]bool, 26)
			cur = 0
		}
	}
	return ans
}
