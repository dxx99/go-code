package main

import (
	"fmt"
)
//todo
type Request struct {
	args []int
	fn func([]int)int
	resultChan chan int
}

func sum(nums []int) int {
	s := 0
	for _, n:= range nums {
		s += n
	}
	return s
}

func main() {
	req := &Request{
		args:       []int{3,4,5},
		fn:         sum,
		resultChan: make(chan int),
	}
	fmt.Println(req)
}
