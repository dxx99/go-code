package main

import "fmt"

func main() {
	nc := Constructor()
	nc.Find(10)
	nc.Change(2, 10)
	nc.Change(1, 10)
	nc.Change(3, 10)
	nc.Change(5, 10)
	nc.Find(10)
	nc.Change(1, 20)
	nc.Find(10)
}


//1.
func bestHand(ranks []int, suits []byte) string {
	color := true
	for i := 1; i < len(suits); i++ {
		if suits[i] != suits[i-1] {
			color = false
		}
	}

	// 同花
	if color {
		return "Flush"
	}

	// 三张
	m := make(map[int]int)
	for i := 0; i < len(ranks); i++ {
		m[ranks[i]]++
	}

	res := "High Card"
	for _, v := range m {
		if v >= 3 {
			return "Three of a Kind"
		}else if v == 2 {
			res = "Pair"
		}
	}
	return res
}

//2.
func zeroFilledSubarray(nums []int) int64 {
	f := func(n int) int64 {
		return int64(n)*int64(n+1)/2
	}

	nums = append(nums, 1)	//添加哨兵
	var res int64
	zeroNum := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			zeroNum++
		}else {
			res += f(zeroNum)
			zeroNum = 0
		}
	}
	return res
}

//3.
type NumberContainers struct {
	bucket map[int]int	// 记录下标->数字的kv值
	vMap map[int][]int   // 记录数字->下标数组值

}


func Constructor() NumberContainers {
	return NumberContainers{
		bucket: make(map[int]int),
		vMap: make(map[int][]int),
	}
}


func (n *NumberContainers) Change(index int, number int)  {
	v, ok := n.bucket[index]
	if ok {	// 删除原来的值
		arr := n.vMap[v]
		fmt.Println(arr)
		for i := 0; i < len(arr); i++ {
			if arr[i] == index {
				n.vMap[v] = append(arr[:i], arr[i+1:]...)
				break
			}
		}
	}

	// 修改bucket值
	n.bucket[index] = number

	// 修改v-keys的数组
	nv, ok := n.vMap[number]
	if ok {
		n.vMap[number] = append(nv, index)
	} else {
		n.vMap[number] = []int{index}
	}
}


func (n *NumberContainers) Find(number int) int {
	fmt.Println(n.vMap)
	v, ok := n.vMap[number]
	if !ok || len(v) == 0 {
		return -1
	}

	min := v[0]
	for i := 1; i < len(v); i++ {
		if v[i] < min {
			min = v[i]
		}
	}
	return min
}

//4.
func shortestSequence(rolls []int, k int) int {
	return 0	
}