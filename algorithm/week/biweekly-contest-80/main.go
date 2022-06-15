package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

func main() {
	fmt.Println(matchReplacement("fool3e7bar", "leet", [][]byte{{'e','3'},{'t','7'},{'t','8'}}))
}

// 1.
func strongPasswordCheckerII(password string) bool {
	if len(password) < 8 {
		return false
	}
	hash := map[byte]bool{
		'!' : true,
		'@' : true,
		'#' : true,
		'$' : true,
		'%' : true,
		'^' : true,
		'&' : true,
		'*' : true,
		'(' : true,
		')' : true,
		'-' : true,
		'+' : true,
	}
	x, d, s, ts := false, false, false, false
	for i := 0; i < len(password); i++ {
		if password[i] >= 'a' && password[i] <= 'z' {
			x = true
		}
		if password[i] >= 'A' && password[i] <= 'Z' {
			d = true
		}
		if password[i] >= '0' && password[i] <= '9' {
			s = true
		}
		if hash[password[i]] {
			ts = true
		}
		if i > 0 && password[i] == password[i-1] {
			return false
		}
	}

	if x && d && s && ts {
		return true
	}
	return false
}

// 2.
func successfulPairs(spells []int, potions []int, success int64) []int {
	ans := make([]int, len(spells))

	sort.Ints(potions)
	for i := 0; i < len(spells); i++ {
		left, right := 0, len(potions)-1
		for left < right {
			mid := (left+right)/2
			if int64(potions[mid]*spells[i]) < success {
				left = mid+1
			}else {
				right = mid
			}
		}
		if int64(potions[left]*spells[i]) >= success {
			ans[i] = len(potions)-left
		}else {
			ans[i] = len(potions)-left-1
		}

	}
	return ans
}

// 3.
func matchReplacement(s string, sub string, mappings [][]byte) bool {

	hashByte := make(map[byte][]string, 0)
	for i := 0; i < len(mappings); i++ {
		if _, ok :=hashByte[mappings[i][0]]; ok {
			hashByte[mappings[i][0]] = append(hashByte[mappings[i][0]], string(mappings[i][1]))
		}else {
			hashByte[mappings[i][0]] = []string{string(mappings[i][0]), string(mappings[i][1])}
		}
	}

	r := ""
	for i := 0; i < len(sub); i++ {
		if v, ok := hashByte[sub[i]]; ok {
			r += "[" + strings.Join(v, "|") + "]"
		}else {
			r += string(sub[i])
		}
	}

	reg := regexp.MustCompile(r)
	return reg.FindString(s) != ""
}



// 4.