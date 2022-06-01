package main

import (
	"container/heap"
	"fmt"
)

func main() {
	obj := Constructor()
	obj.AddNum(1)
	obj.AddNum(2)
	fmt.Println(obj.FindMedian())
	obj.AddNum(3)

	fmt.Println(obj.FindMedian())
	//obj.AddNum(114)
	//fmt.Println(obj.FindMedian())
	//obj.AddNum(132)
	//obj.AddNum(11)
	//obj.AddNum(199)
	//obj.AddNum(189)
	//obj.AddNum(33)
	//fmt.Println(obj.FindMedian())


}

type IntHeap []int

func (h *IntHeap) Len() int {
	return len(*h)
}

func (h *IntHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *IntHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(*h)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type MedianFinder struct {
	large *IntHeap
	small *IntHeap
}

// Constructor initialize your data structure here.
func Constructor() MedianFinder {
	max, min := &IntHeap{}, &IntHeap{}
	heap.Init(max)
	heap.Init(min)
	return MedianFinder{large: max, small: min}
}

func (m *MedianFinder) AddNum(num int)  {
	if m.large.Len() == 0 || num > (*m.large)[0] {
		heap.Push(m.large, num)
	} else {
		heap.Push(m.small, -num)
	}

	// 两栈大小调整
	if m.large.Len() > m.small.Len() +1 {
		heap.Push(m.small, -heap.Pop(m.large).(int))
	}else if m.small.Len() > m.large.Len() + 1 {
		heap.Push(m.large, - heap.Pop(m.small).(int))
	}
}


func (m *MedianFinder) FindMedian() float64 {
	if m.large.Len() < m.small.Len() {
		return float64(-(*m.small)[0])
	} else if m.large.Len() > m.small.Len() {
		return float64((*m.large)[0])
	}else {
		return float64(-(*m.small)[0] + (*m.large)[0])/2
	}
}


/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
