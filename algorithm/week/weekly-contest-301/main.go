package main

import (
	"math"
	"sort"
)

func main() {

}

//1
func fillCups(amount []int) int {
	sort.Ints(amount)
	x, y := amount[1], amount[2]

	for amount[0] > 0 {
		if y > x {
			x++
		}else {
			y++
		}
		amount[0]--
	}
	if x > y {
		return x
	}
	return y
}


//2
type SmallestInfiniteSet struct {
	d map[int]bool
}


func Constructor() SmallestInfiniteSet {
	return SmallestInfiniteSet{make(map[int]bool)}
}


func (s *SmallestInfiniteSet) PopSmallest() int {
	for i := 1; i < math.MaxInt; i++ {
		if !s.d[i] {
			s.d[i] = true
			return i
		}
	}
	return 0
}


func (s *SmallestInfiniteSet) AddBack(num int)  {
	s.d[num] = true
}


//3

//4
