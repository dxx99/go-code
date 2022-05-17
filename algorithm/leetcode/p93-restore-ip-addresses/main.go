package main

import (
	"fmt"
	"strconv"
	"strings"
)

// p93 复原ip地址
// 有效 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。
//
//例如："0.1.2.201" 和 "192.168.1.1" 是 有效 IP 地址，但是 "0.011.255.245"、"192.168.1.312" 和 "192.168@1.1" 是 无效 IP 地址。
//给定一个只包含数字的字符串 s ，用以表示一个 IP 地址，返回所有可能的有效 IP 地址，这些地址可以通过在 s 中插入 '.' 来形成。你 不能 重新排序或删除 s 中的任何数字。你可以按 任何 顺序返回答案。
//
// 
//
//示例 1：
//
//输入：s = "25525511135"
//输出：["255.255.11.135","255.255.111.35"]
//示例 2：
//
//输入：s = "0000"
//输出：["0.0.0.0"]
//示例 3：
//
//输入：s = "101023"
//输出：["1.0.10.23","1.0.102.3","10.1.0.23","10.10.2.3","101.0.2.3"]
// 
//
//提示：
//
//1 <= s.length <= 20
//s 仅由数字组成
//
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/restore-ip-addresses
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	fmt.Println(restoreIpAddresses("25525511135"))
	fmt.Println(restoreIpAddresses("0000"))
	fmt.Println(restoreIpAddresses("101023"))
}

func restoreIpAddresses(s string) []string {
	res := make([]string, 0)
	track := make([]string, 0)

	var backtracking func(start int)
	backtracking = func(start int) {
		if len(track) > 4 {
			return
		}
		// 终止条件
		if start >= len(s) && len(track) == 4 {
			res = append(res, strings.Join(track, "."))
			return
		}

		// 遍历
		for i := start; i < len(s); i++ {
			tmp := s[start:i+1]
			num, _ := strconv.Atoi(tmp)
			if num >= 0 && num <= 255 && tmp == strconv.Itoa(num) {
				track = append(track, tmp)
			}else {
				continue
			}

			// 递归
			backtracking(i+1)
			//回溯
			track = track[:len(track)-1]
		}
	}

	backtracking(0)
	return res
}
