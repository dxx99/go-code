package main

import "fmt"

// 线段树
func main() {
	arr := []int{5,9,7,4,6,1}

	tree := NewLineTree()
	tree.Build(arr)
	fmt.Println(tree.List)
}


type Node struct {
	Val int		// 树节点的值
	Left int	// 左区间
	Right int	// 右区间
}

type LineTree struct {
	Arr []int
	List []Node
}

// NewLineTree 构建线段树
func NewLineTree() *LineTree {
	return &LineTree{Arr: make([]int, 0), List: make([]Node, 0)}
}


// Build 构建线段树
func (l *LineTree) Build(arr []int) {
	idx := 0
	start, end := 0, len(arr)-1

	// 初始化
	l.List = make([]Node, 2*end+3)
	l.Arr = arr

	l.build(arr, idx, start, end)
}

// n int 树的索引
func (l *LineTree) build(arr []int, n int, start int, end int) {
	// 终止条件
	if start == end {
		l.List[n] = Node{
			Val:   arr[start],
			Left:  start,
			Right: end,
		}
		return
	}

	mid := (start + end) >> 1
	leftNode := 2 * n + 1
	rightNode := 2 * n + 2

	l.build(arr, leftNode, start, mid)
	l.build(arr, rightNode, mid+1, end)

	// 设置区间的值
	l.List[n] = Node{
		Val:   l.List[leftNode].Val + l.List[rightNode].Val,
		Left:  start,
		Right: end,
	}
}

// Update 更新数据
func (l *LineTree) Update(idx int , val int) {
	start, end := 0, len(l.Arr)-1
	l.updateTree(0, start, end, idx, val)
}

func (l *LineTree) updateTree(n int, start int, end int, idx int, val int) {
	if start == end {
		l.Arr[idx] = val
		l.List[idx].Val = val
		return
	}

	mid := (start + end) >> 1
	leftNode := 2 * n + 1
	rightNode := 2 * n + 2
	if idx >= start && idx <= mid {
		l.updateTree(leftNode, start, mid, idx, val)
	} else {
		l.updateTree(rightNode, mid+1, end, idx, val)
	}

	// 回溯更新父节点的和
	l.List[n].Val = l.List[leftNode].Val + l.List[rightNode].Val
}


func (l *LineTree) Query(left int, right int) int {
	start, end := 0, len(l.Arr)-1
	return l.queryTree(0, start, end, left, right)
}

func (l *LineTree) queryTree(n int, start int, end int, left int, right int) int {

	// 终止条件,查到叶子节点了，因为只有一个元素则，直接返回
	if start == end {
		return l.List[n].Val
	}
	// 其他条件终止，不在区间中
	if end < left || start > right {
		return 0
	}
	// 完全包含在区间中，不需要额外的查找
	if start>= left && end <= right {
		return l.List[n].Val
	}
	
	mid := (start + end) >> 1
	leftNode := 2*n+1
	rightNode := 2*n+2
	leftSum := l.queryTree(leftNode, start, mid, left, right)
	rightSum := l.queryTree(rightNode, mid+1, end, left, right)
	return leftSum+rightSum
}