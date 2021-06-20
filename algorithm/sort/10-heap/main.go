package main

import (
	"container/heap"
	"fmt"
	"math/rand"
	"time"
)

// 生成随机数组
func getRandArray(n int) []int {
	rand.Seed(time.Now().UnixNano())

	var res []int
	for i := 0; i < n; i++ {
		res = append(res, rand.Intn(1e3))
	}
	return res
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func heapSort(nums []int) []int {
	h := new(IntHeap)
	heap.Init(h)
	for _, item := range nums {
		heap.Push(h, item)
	}
	k := 0
	for h.Len() > 0 {
		nums[k] = heap.Pop(h).(int)
		k++
	}
	return nums
}

// 堆排序：
func main() {
	nums := getRandArray(10)
	fmt.Println(nums)
	fmt.Println(heapSort(nums))
}
