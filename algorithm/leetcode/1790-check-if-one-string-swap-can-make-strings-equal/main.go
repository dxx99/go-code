package main

import "fmt"

func main() {
	//fmt.Println(areAlmostEqual("bank", "kanb"))
	//fmt.Println(areAlmostEqual("attack", "defend"))
	//fmt.Println(areAlmostEqual("kelb", "kelb"))
	//fmt.Println(areAlmostEqual("abcd", "dcba"))

	fmt.Println(areAlmostEqual("qgqeg", "gqgeq"))
}

func areAlmostEqual(s1 string, s2 string) bool {
	x, num := -1, 0
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			num++
			if num >= 3 {
				return false
			}
			if x == -1 {
				x = -1
			} else {
				if s1[x] != s2[i] || s1[i] != s2[x] {
					return false
				}
			}
		}
	}
	return num%2 == 0
}
