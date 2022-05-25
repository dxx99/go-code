package main

import "fmt"

// 输入两个整数序列，第一个序列表示栈的压入顺序，请判断第二个序列是否为该栈的弹出顺序。假设压入栈的所有数字均不相等。例如，序列 {1,2,3,4,5} 是某栈的压栈序列，序列 {4,5,3,2,1} 是该压栈序列对应的一个弹出序列，但 {4,3,5,1,2} 就不可能是该压栈序列的弹出序列。
//
// 
//
//示例 1：
//
//输入：pushed = [1,2,3,4,5], popped = [4,5,3,2,1]
//输出：true
//解释：我们可以按以下顺序执行：
//push(1), push(2), push(3), push(4), pop() -> 4,
//push(5), pop() -> 5, pop() -> 3, pop() -> 2, pop() -> 1
//示例 2：
//
//输入：pushed = [1,2,3,4,5], popped = [4,3,5,1,2]
//输出：false
//解释：1 不能在 2 之前弹出。
// 
//
//提示：
//
//0 <= pushed.length == popped.length <= 1000
//0 <= pushed[i], popped[i] < 1000
//pushed 是 popped 的排列。
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/zhan-de-ya-ru-dan-chu-xu-lie-lcof
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	fmt.Println(validateStackSequences([]int{1,1,3,4,5}, []int{4,5,3,1,1}))
}

func validateStackSequencesV2(pushed []int, popped []int) bool {
	if len(pushed) == 0 && len(popped) == 0 {
		return true
	}
	stack := make([]int, 0)
	pk := 0
	for i := 0; i < len(pushed); i++ {
		// 先入栈
		stack = append(stack, pushed[i])

		// 再比较栈顶元素，如果相等则出栈
		for len(stack) > 0 {
			// 终止条件
			if pk == len(popped)-1 {
				return true
			}

			top := stack[len(stack)-1]
			if top != popped[pk] {
				break
			}
			stack = stack[:len(stack)-1]
			pk++
		}
	}

	return false
}

// 题示： 简化了逻辑
// 1. 0 <= pushed.length == popped.length <= 1000
// 2. 0 <= pushed[i], popped[i] < 1000
// 3. pushed 是 popped 的排列。
func validateStackSequences(pushed []int, popped []int) bool {
	stack := make([]int, 0)
	var idx int

	for _,v :=range pushed {
		stack = append(stack, v)

		for len(stack) > 0 && stack[len(stack)-1] == popped[idx] {
			stack = stack[:len(stack)-1] // poped the last element
			idx++
		}
	}

	return len(stack) == 0

}
