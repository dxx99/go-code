package main

import (
	"fmt"
	"sort"
)

// p406 根据身高重建队列
// 假设有打乱顺序的一群人站成一个队列，数组 people 表示队列中一些人的属性（不一定按顺序）。每个 people[i] = [hi, ki] 表示第 i 个人的身高为 hi ，前面 正好 有 ki 个身高大于或等于 hi 的人。
//
//请你重新构造并返回输入数组 people 所表示的队列。返回的队列应该格式化为数组 queue ，其中 queue[j] = [hj, kj] 是队列中第 j 个人的属性（queue[0] 是排在队列前面的人）。
//
// 
//
//示例 1：
//
//输入：people = [[7,0],[4,4],[7,1],[5,0],[6,1],[5,2]]
//输出：[[5,0],[7,0],[5,2],[6,1],[4,4],[7,1]]
//解释：
//编号为 0 的人身高为 5 ，没有身高更高或者相同的人排在他前面。
//编号为 1 的人身高为 7 ，没有身高更高或者相同的人排在他前面。
//编号为 2 的人身高为 5 ，有 2 个身高更高或者相同的人排在他前面，即编号为 0 和 1 的人。
//编号为 3 的人身高为 6 ，有 1 个身高更高或者相同的人排在他前面，即编号为 1 的人。
//编号为 4 的人身高为 4 ，有 4 个身高更高或者相同的人排在他前面，即编号为 0、1、2、3 的人。
//编号为 5 的人身高为 7 ，有 1 个身高更高或者相同的人排在他前面，即编号为 1 的人。
//因此 [[5,0],[7,0],[5,2],[6,1],[4,4],[7,1]] 是重新构造后的队列。
//示例 2：
//
//输入：people = [[6,0],[5,0],[4,0],[3,2],[2,2],[1,4]]
//输出：[[4,0],[5,0],[2,2],[3,2],[1,4],[6,0]]
// 
//
//提示：
//
//1 <= people.length <= 2000
//0 <= hi <= 106
//0 <= ki < people.length
//题目数据确保队列可以被重建
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/queue-reconstruction-by-height
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	fmt.Println(reconstructQueue([][]int{{7,0},{4,4},{7,1},{5,0},{6,1},{5,2}}))
	fmt.Println(reconstructQueue([][]int{{6,0},{5,0},{4,0},{3,2},{2,2},{1,4}}))

	fmt.Println("V2.........")
	fmt.Println(reconstructQueueV2([][]int{{7,0},{4,4},{7,1},{5,0},{6,1},{5,2}}))
	fmt.Println(reconstructQueueV2([][]int{{6,0},{5,0},{4,0},{3,2},{2,2},{1,4}}))
}


func reconstructQueue(people [][]int) [][]int {
	// 先排序
	// 规则： 身高高的人排在前面，身高相等的k值小的在前面
	// 这样就可以得到一个规则，前面的节点一定比本节点高，然后就可以用插入排序，将后面的元素插入进去
	// 局部最优：优先按身高高的人的k来插入，插入操作过后的人，满足队列属性
	// 全局最优：最后做完插入操作，整个队列满足题目队列属性
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}
		return people[i][0] > people[j][0]
	})
	for i := 0; i < len(people); i++ {
		k := people[i][1]
		if k != i {
			people = append(append(people[:k], append([][]int{people[i]}, people[k:i]...)...), people[i+1:]...)
		}
	}

	return people
}

func reconstructQueueV2(people [][]int) [][]int {
	// 先排序
	// 规则： 身高高的人排在前面，身高相等的k值小的在前面
	// 这样就可以得到一个规则，前面的节点一定比本节点高，然后就可以用插入排序，将后面的元素插入进去
	// 局部最优：优先按身高高的人的k来插入，插入操作过后的人，满足队列属性
	// 全局最优：最后做完插入操作，整个队列满足题目队列属性
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}
		return people[i][0] > people[j][0]
	})

	ans := make([][]int, 0)
	for i := 0; i < len(people); i++ {
		k := people[i][1]
		ans = append(ans[:k], append([][]int{people[i]}, ans[k:]...)...)
	}

	return ans
}
