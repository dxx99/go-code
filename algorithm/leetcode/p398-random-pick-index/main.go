package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	fmt.Println(rand.Intn(3))
}


type Solution struct {
	m map[int][]int
}


func Constructor(nums []int) Solution {
	m := make(map[int][]int)
	for i := 0; i < len(nums); i++ {
		if v, ok := m[nums[i]]; ok {
			m[nums[i]] = append(v, i)
		} else {
			m[nums[i]] = append([]int{}, i)
		}
	}
	return Solution{m: m}
}


func (this *Solution) Pick(target int) int {
	v, ok := this.m[target]
	if !ok {
		return -1
	}
	rand.Seed(time.Now().Unix())
	return v[rand.Intn(len(v))]
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Pick(target);
 */