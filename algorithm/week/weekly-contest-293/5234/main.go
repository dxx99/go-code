package main

import "fmt"

func main() {
	fmt.Println(removeAnagrams([]string{"abba","baba","bbaa","cd","cd"}))
	fmt.Println(removeAnagrams([]string{"a","b","c","d","e"}))
}

func removeAnagrams(words []string) []string {
	res := make([]string, 0)

	last := [26]int{}
	for i := 0; i < len(words); i++ {
		cur := [26]int{}
		for j := 0; j < len(words[i]); j++ {
			cur[words[i][j] - 'a']++
		}
		if last != cur {
			res = append(res, words[i])
		}
		last = cur
	}

	return res
}
