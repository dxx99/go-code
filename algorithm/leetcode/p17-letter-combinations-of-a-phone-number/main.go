package main

import "fmt"

// p17 电话号码的字母组合
// 给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。
//
//给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
//
//
//
// 
//
//示例 1：
//
//输入：digits = "23"
//输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]
//示例 2：
//
//输入：digits = ""
//输出：[]
//示例 3：
//
//输入：digits = "2"
//输出：["a","b","c"]
// 
//
//提示：
//
//0 <= digits.length <= 4
//digits[i] 是范围 ['2', '9'] 的一个数字。
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/letter-combinations-of-a-phone-number
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	fmt.Println(letterCombinations("23"))
	fmt.Println(letterCombinations(""))
	fmt.Println(letterCombinations("2"))
}


func letterCombinations(digits string) []string {
	res := make([]string, 0)  // 存放找到的结果
	cur := ""
	if len(digits) == 0 {
		return res
	}

	letterMap := []string{
		"",	//0
		"",	//1
		"abc",
		"def",
		"ghi",
		"jkl",
		"mno",
		"pqrs",
		"tuv",
		"wxyz",
	}

	var backtracking func(l int, start int)
	backtracking = func(l, start int) {
		// 成功终止条件
		if len(cur) == l {
			res = append(res, cur)
			return
		}
		
		//遍历
		letters := letterMap[digits[start] - '0']	// 表示abc, def等
		for j := 0; j < len(letters); j++ {
			cur += string(letters[j])
			// 递归
			backtracking(l, start+1)

			// 回溯, 清掉上一个元素
			cur = cur[:len(cur)-1]
		}
	}

	backtracking(len(digits), 0)
	return res
}
