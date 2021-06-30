package main

import "fmt"

// 切片基本操作
func main() {
	s := make([]int, 0)

	// 增加
	s = append(s, []int{1, 2, 3, 4}...)
	fmt.Printf("after append item, slice = %v\n", s)

	// 合并slice
	b := []int{3, 4, 5, 6}
	s = append(s, b...)
	fmt.Printf("after merge slice, slice = %v\n", s)

	// 在指定位置插入
	k, v := 2, 44
	s = append(s[:k], append([]int{v}, s[k:]...)...)
	fmt.Printf("at k insert, slice = %v\n", s)

	// pop， 弹出最后的元素
	x, s := s[len(s)-1], s[:len(s)-1]
	fmt.Printf("last item is %d\n", x)

	// push, 向后压入一个元素
	w := 55
	s = append(s, w)
	fmt.Printf("after push, slice = %v\n", s)

	// shift 弹出第一个元素
	f, s := s[0], s[1:]
	fmt.Printf("after shift, item = %d, slice = %v\n", f, s)

	// unshift 向队首压入一个元素
	z := 33
	s = append([]int{z}, s...)
	fmt.Printf("slice is %v\n", s)

	// reversing 数组元素反转
	for i := len(s)/2 - 1; i >= 0; i-- {
		k := len(s) - 1 - i
		s[i], s[k] = s[k], s[i]
	}
	fmt.Printf("reverse slice is = %v\n", s)
}
