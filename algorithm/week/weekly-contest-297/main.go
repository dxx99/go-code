package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(distributeCookies([]int{8,15,10,20,8},2))
	fmt.Println(distributeCookies([]int{6,1,3,2,2,4,1,2},3))
	fmt.Println(distributeCookies([]int{1,8,16,5,6,14},6))

}

//1.
func calculateTax(brackets [][]int, income int) float64 {
	tax := 0.0
	for i := 0; i < len(brackets); i++ {
		if income >= brackets[i][0] {
			if i == 0  {
				tax += float64(brackets[i][0] * brackets[i][1])/100
			}else {
				tax += float64((brackets[i][0] - brackets[i-1][0]) * brackets[i][1])/100
			}
		}else {
			if i == 0 {
				tax += float64(income*brackets[i][1])/100
			}else {
				tax += float64((income - brackets[i-1][0]) * brackets[i][1])/100
			}
			break
		}
	}
	return tax
}


// 3.
func distributeCookies(cookies []int, k int) int {
	sort.Ints(cookies)
	fmt.Println(cookies)

	ans := make([]int, k)
	flag := false
	for i := 0; i < len(cookies); i++ {
		index := i % k
		if index ==0  {
			flag = !flag
		}
		if flag {
			ans[index] += cookies[i]
		}else {
			ans[k-index-1] += cookies[i]
		}
	}
	max := 0
	for i := 0; i < len(ans); i++ {
		if ans[i] > max {
			max = ans[i]
		}
	}
	return max
}

//TODO
func subset(cookies []int, l int, s int) bool {
	if s == 0 {
		return true
	}else if l == 0 {
		return cookies[0] == s
	}else if cookies[l]>s {
		return subset(cookies, l-1, s)
	}else {
		a := subset(cookies, l-1, s-cookies[l])
		b := subset(cookies, l-1, s)
		return a || b
	}
}