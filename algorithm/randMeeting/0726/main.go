package main

import "fmt"

func main() {
	//fmt.Println(highestPeak([][]int{{0,1},{0,0}}))
	//fmt.Println(asteroidCollision([]int{5,10,-5}))

	fmt.Println(checkPartitioning("abcbdd"))
	fmt.Println(checkPartitioning("bcbddxy"))
}

//1
func highestPeak(isWater [][]int) [][]int {
	queue := make([][]int, 0)

	for i := 0; i < len(isWater); i++ {
		for j := 0; j < len(isWater[i]); j++ {
			if isWater[i][j] == 1 {
				isWater[i][j] = 0
				queue = append(queue, []int{i, j})
			}else {
				isWater[i][j] = -1
			}
		}
	}
	m, n := len(isWater), len(isWater[0])
	xl, yl := []int{1,0,0,-1}, []int{0,1,-1,0}

	// 取队列数据
	for len(queue) > 0 {
		tmp := queue
		queue = make([][]int, 0)
		for i := 0; i < len(tmp); i++ {
			x, y := tmp[i][0], tmp[i][1]
			ov := isWater[x][y]
			for k := 0; k < 4; k++ {
				cx, cy := x+xl[k], y+yl[k]
				if cx >= 0  && cx < m && cy >= 0 && cy < n && isWater[cx][cy] == -1 {
					isWater[cx][cy] = ov+1
					queue = append(queue, []int{cx, cy})
				}
			}
		}
	}

	return isWater
}

//2
func asteroidCollision(asteroids []int) []int {
	stack := make([]int, 0)

	for _, aster := range asteroids {
		alive := true
		// for循环，比较当前行星与栈顶元素
		for alive && aster < 0 && len(stack) > 0 && stack[len(stack)-1] > 0  {

			// 判断当前的aster是否还存在
			alive = stack[len(stack)-1] < -aster

			// 出栈，当前碰撞的行星比栈顶的行星质量大
			if stack[len(stack)-1] <= -aster {
				stack = stack[:len(stack)-1]
			}
		}

		// 当前行星是否还存在
		if alive {
			stack = append(stack, aster)
		}
	}

	return stack
}

//3
func checkPartitioning(s string) bool {
	f := func(str string) bool {
		if len(str) == 0  {
			return false
		}
		left, right := 0, len(str)-1
		for left < right {
			if str[left] == str[right] {
				left++
				right--
			}else {
				return false
			}
		}
		return true
	}

	for i := 1; i < len(s)-1; i++ {
		if f(s[:i]) {
			for j := i+1; j < len(s); j++ {
				//fmt.Println(s[:i], s[i:j], s[j:])
				if f(s[i:j]) && f(s[j:]) {
					return true
				}
			}
		}
	}

	return false
}

