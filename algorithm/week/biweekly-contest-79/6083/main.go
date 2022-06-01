package main

import "fmt"

func main() {
	fmt.Println(digitCount("1210"))
	fmt.Println(digitCount("030"))
	fmt.Println(digitCount(""))
}

func digitCount(num string) bool {
	arr := [10]int{}

	for i := 0; i < len(num); i++ {
		arr[num[i] - '0']++
	}

	for i := 0; i < len(num); i++ {
		v := num[i] - '0'
		if arr[i] != int( v) {
			return false
		}
	}
	return true
}
