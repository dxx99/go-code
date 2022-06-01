package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	h := new(MyHeap)
	heap.Init(h)

	h.Push(11)
	h.Push(22)
	h.Push(2)
	h.Push(4)
	h.Push(5)
	fmt.Println(h)
	fmt.Println(h.Pop())
	fmt.Println(h.Pop())

}

type MyHeap []int

func (m *MyHeap) Len() int {
	return len(*m)
}

func (m *MyHeap) Less(i, j int) bool {
	return (*m)[i] < (*m)[j]
}

func (m *MyHeap) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func (m *MyHeap) Push(x interface{}) {
	*m = append(*m, x.(int))
}

func (m *MyHeap) Pop() (v interface{}) {
	*m, v = (*m)[:m.Len()-1], (*m)[m.Len()-1]
	return
}

func getLeastNumbers(arr []int, k int) []int {
	sort.Ints(arr)
	return arr[:k]
}
