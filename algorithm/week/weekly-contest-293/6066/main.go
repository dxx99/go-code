package main

import "fmt"

//["CountIntervals","count","add","count"]
//[[],[],[1,1000000000],[]]
func main() {
	obj := Constructor()
	fmt.Println(obj.Count())
	obj.Add(1,1000000000)

	obj.Add(2,3)
	obj.Add(7,10)
	fmt.Println(obj.Count())
	fmt.Println(obj.Count())


}


type CountIntervals struct {
	s [][]int
	m map[int]int
	c int
}


func Constructor() CountIntervals {
	return CountIntervals{
		s: make([][]int, 0),
		m: make(map[int]int),
		c: 0,
	}
}


func (this *CountIntervals) Add(left int, right int)  {

	for _, item := range this.s {
		iLeft, iRight := item[0], item[1]



		// 有重合
		if right > iLeft && right < iRight && left < iLeft {
			this.c += iLeft - left + 1
		}

	}
	
	this.s = append(this.s, []int{left, right})
	this.m[left] = right - left + 1
	this.m[right] = right - right + 1
	
	
}


func (this *CountIntervals) Count() int {
	return this.c
}


/**
 * Your CountIntervals object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(left,right);
 * param_2 := obj.Count();
 */
