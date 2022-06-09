package main

import "fmt"

//给定一个整数数组 temperatures ，表示每天的温度，返回一个数组 answer ，其中 answer[i] 是指在第 i 天之后，才会有更高的温度。如果气温在这之后都不会升高，请在该位置用 0 来代替。
//
// 
//
//示例 1:
//
//输入: temperatures = [73,74,75,71,69,72,76,73]
//输出: [1,1,4,2,1,1,0,0]
//示例 2:
//
//输入: temperatures = [30,40,50,60]
//输出: [1,1,1,0]
//示例 3:
//
//输入: temperatures = [30,60,90]
//输出: [1,1,0]
// 
//
//提示：
//
//1 <= temperatures.length <= 105
//30 <= temperatures[i] <= 100
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/daily-temperatures
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	fmt.Println(dailyTemperatures([]int{73,74,75,71,69,72,76,73}))	// [1 1 4 2 1 1 0 0]
	fmt.Println(dailyTemperatures([]int{30,40,50,60}))	// [1 1 1 0]

}

func dailyTemperatures(temperatures []int) []int {
	ans := make([]int, len(temperatures))
	stack := make([]int, 0)

	for i := 0; i < len(temperatures); i++ {
		for len(stack) != 0 && temperatures[stack[len(stack)-1]] < temperatures[i] {
			// 统一出栈的时候对记录的数据做处理
			top := stack[len(stack)-1]
			ans[top] = i - top

			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return ans
}
