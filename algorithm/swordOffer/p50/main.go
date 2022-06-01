package main

import "fmt"

func main() {
	fmt.Println(firstUniqChar("abaccdeff"))
}

func firstUniqChar(s string) byte {
	hashByte := [26]int{}

	for i := 0; i < len(s); i++ {
		hashByte[s[i] - 'a']++
	}

	for i := 0; i < len(s); i++ {
		if hashByte[s[i] - 'a'] == 1 {
			return s[i]
		}
	}
	return ' '
}
