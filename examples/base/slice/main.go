package main

import "fmt"

func main() {
	//#1 切片传参【传值，传址】
	s := []int{1,2,3}
	newS := myAppend(s)
	fmt.Println(s, newS)	// output: [1 2 3] [1 2 3 100]
	s = newS
	myAppendPtr(&s)
	fmt.Println(s)	// output: [1 2 3 100 100]


	//#2 因为底层的数组被修改了
	s2 := []int{1,2,3,4,5}
	modify(s2)
	fmt.Printf("%v\n", s2)	// output: [1 11 3 4 5]

	//#3 切片容量问题
	demo1()	// output: len=5, cap=6
	demo2() // output: len=5, cap=8
}

// 直接会修改掉底层的数组
func modify(s []int) {
	s[1] = 11
}


// 切片作为参数的传参例子
func myAppend(s []int) []int {
	s = append(s, 100)
	return s
}
func myAppendPtr(s *[]int)  {
	*s = append(*s, 100)
	return
}


// output: len=5, cap=6
// 为啥不是 len=5, cap=8, 如果按正常的扩容的逻辑，小于1024就是，double扩容
// roundupsize() 会做一些内存对齐操作
func demo1()  {
	s := []int{1,2}
	s = append(s,4,5,6)
	fmt.Printf("len=%d, cap=%d\n",len(s),cap(s))
}


// output: len=5, cap=8
func demo2()  {
	s := []int{1,2}
	for i := 4; i < 7; i++ {
		s = append(s, i)
	}
	fmt.Printf("len=%d, cap=%d\n",len(s),cap(s))
}
