package main

import "fmt"

func main() {
	obj := Constructor(2,5)
	fmt.Println(obj.Gather(4,0))
	fmt.Println(obj.Gather(3,0))
	fmt.Println(obj.Scatter(5,1))
	fmt.Println(obj.Scatter(5,1))
}

type BookMyShow struct {
	bucket [][]int
}


func Constructor(n int, m int) BookMyShow {
	b := make([][]int, n)
	for i := 0; i < len(b); i++ {
		b[i] = make([]int, m)
	}
	return BookMyShow{bucket: b}
}


func (b *BookMyShow) Gather(k int, maxRow int) []int {
	s := -1 // 开始标记点
	needChangePoint := make([][]int, 0)
	for i := 0; i <= maxRow; i++ {
		for j := 0; j < len(b.bucket[i]); j++ {

			if b.bucket[i][j] != 1 && s == -1 {
				s = 1
				needChangePoint = append(needChangePoint, []int{i,j})
				bNum := k
				bNum--
				for z := j+1; z < len(b.bucket[i]); z++ {
					if bNum == 0 {
						for _, p := range needChangePoint {
							b.bucket[p[0]][p[1]] = 1
						}
						return needChangePoint[0]
					}
					if b.bucket[i][z] != 1 {
						needChangePoint = append(needChangePoint, []int{i,z})
					}else {
						needChangePoint = [][]int{}
						s = -1
						j = z
						break
					}
					bNum--
				}
				s = -1
				continue
			}
		}
	}

	return []int{}
}


func (b *BookMyShow) Scatter(k int, maxRow int) bool {
	needChangePoint := make([][]int, 0)
	bNum := k
	for i := 0; i <= maxRow; i++ {
		for j := 0; j < len(b.bucket[i]); j++ {
			if bNum == 0 {
				for _, p := range needChangePoint {
					b.bucket[p[0]][p[1]] = 1
				}
				return true
			}
			if b.bucket[i][j] != 1 {
				needChangePoint = append(needChangePoint, []int{i,j})
				bNum--
			}
		}
	}

	return  false
}


/**
 * Your BookMyShow object will be instantiated and called as such:
 * obj := Constructor(n, m);
 * param_1 := obj.Gather(k,maxRow);
 * param_2 := obj.Scatter(k,maxRow);
 */
