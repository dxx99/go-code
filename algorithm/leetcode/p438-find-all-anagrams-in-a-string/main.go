package main

import "fmt"

func main() {
	fmt.Println(findAnagrams("cbaebabacd", "abc"))
	fmt.Println(findAnagrams("abab", "ab"))
	fmt.Println(findAnagrams("aaaaaaaaaa", "aaaaaaaaaaaaa"))
}

func findAnagrams(s string, p string) []int {
	ans := make([]int, 0)
	if len(p) > len(s) {
		return ans
	}

	p1 := [26]int{}
	for i := 0; i < len(p); i++ {
		p1[p[i]-'a']++
	}

	p2 := [26]int{}
	for i := 0; i < len(p)-1; i++ {
		p2[s[i]-'a']++
	}

	for i := len(p)-1; i < len(s); i++ {
		p2[s[i]-'a']++
		if p2 == p1 {
			ans = append(ans, i-len(p)+1)
		}
		p2[s[i-len(p)+1]-'a']--
	}
	return ans
}
