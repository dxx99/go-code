package main

import "fmt"

func main() {
	fmt.Println(isLongPressedName("alexd", "ale"))
}

//1
// "alexd"
//"ale"
func isLongPressedName(name string, typed string) bool {
	k := 0
	last := name[0]
	for i := 0; i < len(typed); i++ {
		if k < len(name) && name[k] == typed[i] {
			last = name[k]
			k++
			continue
		}

		// 不是重复的值就不行
		if last != typed[i] {
			return false
		}
	}
	if k != len(name) {
		return false
	}
	return true
}

//2
func numberOfBoomerangs(points [][]int) int {
	return 0
}

//3
