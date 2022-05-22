package main

import "fmt"

func main() {
	fmt.Println(percentageLetter("foobar", 'o'))
	fmt.Println(percentageLetter("jjjj", 'k'))
}

func percentageLetter(s string, letter byte) int {
	c := 0
	for i := 0; i < len(s); i++ {
		if s[i] == letter {
			c++
		}
	}
	return c*100/len(s)
}