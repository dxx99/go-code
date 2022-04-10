package main

import (
	"bytes"
	"fmt"
)

// p6 Z字形变换
// 将一个给定字符串 s 根据给定的行数 numRows ，以从上往下、从左到右进行 Z 字形排列。
//
//比如输入字符串为 "PAYPALISHIRING" 行数为 3 时，排列如下：
//
//P   A   H   N
//A P L S I I G
//Y   I   R
//之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："PAHNAPLSIIGYIR"。
//
//请你实现这个将字符串进行指定行数变换的函数：
//
//string convert(string s, int numRows);
// 
//
//示例 1：
//
//输入：s = "PAYPALISHIRING", numRows = 3
//输出："PAHNAPLSIIGYIR"
//示例 2：
//输入：s = "PAYPALISHIRING", numRows = 4
//输出："PINALSIGYAHRPI"
//解释：
//P     I    N
//A   L S  I G
//Y A   H R
//P     I
//示例 3：
//
//输入：s = "A", numRows = 1
//输出："A"
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/zigzag-conversion
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	fmt.Println(convert("PAYPALISHIRING", 3) == "PAHNAPLSIIGYIR")	//
	fmt.Println(convert("PAYPALISHIRING", 4) == "PINALSIGYAHRPI")	//
	fmt.Println("V2...............")
	fmt.Println(convertV2("PAYPALISHIRING", 3) == "PAHNAPLSIIGYIR")	//
	fmt.Println(convertV2("PAYPALISHIRING", 4) == "PINALSIGYAHRPI")	//
}

// 利用二维矩阵请求
func convert(s string, numRows int) string {
	l, r := len(s), numRows

	// 处理特殊情况
	if r == 1 || r > l {
		return s
	}

	// 执行一轮需要的次数，方便取模
	t := 2*r - 2
	// c 的计算的 周期{(l/t) + 1} * 每个周期列的长度(r-1)
	// 需要向上取整
	c := (l/t +1) * (r-1)
	arr := make([][]byte, r)
	for i := range arr {
		arr[i] = make([]byte, c)
	}

	x, y :=0, 0
	for i, ch := range s {
		arr[x][y] = byte(ch)
		if i%t < r-1 {
			x++	// 向下移动
		} else {
			x--
			y++	// 向右上移动
		}
	}

	ans := make([]byte, 0)
	for _, row := range arr {
		for _, b := range row {
			if b > 0 {
				ans = append(ans, b)
			}
		}
	}

	return string(ans)
}

// 压缩数组
func convertV2(s string, numRows int) string {
	r := numRows
	if r == 1 || r >= len(s) {
		return s
	}

	arr := make([][]byte, r)
	t, x := r*2-2, 0	// 每个周期的元素，与每层的下标
	for i, ch := range s {
		arr[x] = append(arr[x], byte(ch))

		// 进行层的切换，其中r-1为拐角点
		if i%t < r-1 {
			x++
		}else {
			x--
		}
	}

	return string(bytes.Join(arr, nil))
}

func convertV3(s string, numRows int) string {
	l, r := len(s), numRows
	if r == 1 || r >= l {
		return s
	}

	t := 2*r - 2	// 一个周期最多存在的元素个数
	ans := make([]byte, 0)
	for i := 0; i < r; i++ {	// 枚举矩阵的行
		for j := 0; j + i < l; j += t {		// 枚举每个周期的起始下标
			ans = append(ans, s[j+i])	// 当前周期的第一个字符
			if 0 <i && i < r-1 && j+t-i < l {
				ans = append(ans, s[j+t-i]) // 当前周期的第二个字符， 一个周期最多只能有两个元素
			}
			
		}
	}
	return string(ans)
}
