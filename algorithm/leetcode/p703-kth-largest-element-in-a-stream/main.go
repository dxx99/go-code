package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	//obj := Constructor(3, []int{5,4,8,10})
	//fmt.Println(obj.Add(3))
	demo()
}

func demo() {
	kd := &KthLargest{
		k:        0,
	}

	heap.Init(kd)

	heap.Push(kd, 1)
	heap.Push(kd, 3)
	heap.Push(kd, 7)
	heap.Push(kd, 6)

	fmt.Println("k num is ", kd.k)
	fmt.Println(heap.Pop(kd))
	fmt.Println(heap.Pop(kd))
	fmt.Println(heap.Pop(kd))
	fmt.Println(heap.Pop(kd))
}


// KthLargest 小顶堆实现
type KthLargest struct {
	sort.IntSlice
	k int
}

func (k *KthLargest) Push(x interface{}) {
	k.IntSlice = append(k.IntSlice, x.(int))
	k.k++
}

func (k *KthLargest) Pop() interface{} {
	x := k.IntSlice[len(k.IntSlice)-1]
	k.IntSlice = k.IntSlice[:len(k.IntSlice)-1]
	k.k--
	return x
}

func Constructor(k int, nums []int) KthLargest {
	kd := KthLargest{
		k:        k,
	}
	for _, num := range nums {
		kd.Add(num)
	}
	return kd
}

// Add 插入一个元素，返回第k大的元素
// 堆中只存储三个元素
func (k *KthLargest) Add(val int) int {
	heap.Push(k, val)
	if k.Len()-1 > k.k {
		heap.Pop(k)
	}
	return k.IntSlice[0]
}