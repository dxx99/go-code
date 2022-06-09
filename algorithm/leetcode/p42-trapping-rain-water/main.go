package main

import "fmt"

func main() {
	fmt.Println(trap([]int{0,1,0,2,1,0,1,3,2,1,2,1}))
	fmt.Println(trap([]int{4,2,0,3,2,5}))
	fmt.Println(trapV2([]int{0,1,0,2,1,0,1,3,2,1,2,1}))
	fmt.Println(trapV2([]int{4,2,0,3,2,5}))
	fmt.Println(trapV3([]int{0,1,0,2,1,0,1,3,2,1,2,1}))
	fmt.Println(trapV3([]int{4,2,0,3,2,5}))
}

// 动态规划求解
func trapV3(height []int) int {
	if len(height) <= 2 {
		return 0
	}

	maxLeft := make([]int, len(height))
	maxRight := make([]int, len(height))

	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}

	// 得到该点左右两边的最高点
	maxLeft[0] = height[0]
	for i := 1; i < len(height); i++ {
		maxLeft[i] = max(height[i], maxLeft[i-1])
	}
	maxRight[len(height)-1] = height[len(height)-1]
	for i := len(height)-2; i >= 0 ; i-- {
		maxRight[i] = max(height[i], maxRight[i+1])
	}

	// 求和
	ans := 0
	for i := 0; i < len(height); i++ {
		if min(maxLeft[i], maxRight[i]) - height[i] > 0 {
			ans += min(maxLeft[i], maxRight[i]) - height[i]
		}
	}

	return ans
}

// 单调栈
func trapV2(height []int) int {
	ans := 0
	if len(height) == 0 {
		return ans
	}

	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}

	stack := make([]int, 0)
	stack = append(stack, 0)
	for i := 1; i < len(height); i++ {

		for len(stack) != 0 && height[i] > height[stack[len(stack)-1]] {
			mid := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) != 0 {	// 说明栈至少有2个元素，可以形成一个凹槽
				h := min(height[stack[len(stack)-1]], height[i]) - height[mid]
				w := i - stack[len(stack)-1] -1
				ans += h*w
			}
		}

		// 将当前元素入栈
		stack = append(stack, i)
	}
	return ans
}

func trap(height []int) int {
	ans := 0
	if len(height) == 0 {
		return ans
	}

	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}

	stack := make([]int, 0)
	stack = append(stack, 0)
	for i := 1; i < len(height); i++ {
		if height[stack[len(stack)-1]] > height[i] {
			stack = append(stack, i)
		}else if height[stack[len(stack)-1]] == height[i]  {
			stack = append(stack, i)
		}else{
			for len(stack) != 0 && height[i] > height[stack[len(stack)-1]] {
				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if len(stack) != 0 {
					h := min(height[stack[len(stack)-1]], height[i]) - height[top]
					w := i - stack[len(stack)-1] -1
					ans += h*w
				}
			}
			stack = append(stack, i)
		}
	}
	return ans
}
