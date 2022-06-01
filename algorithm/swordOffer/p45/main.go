package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)
// 若拼接字符串 x + y > y + xx+y>y+x ，则 xx “大于” yy ；
//反之，若 x + y < y + xx+y<y+x ，则 xx “小于” yy ；
//
//作者：jyd
//链接：https://leetcode.cn/problems/ba-shu-zu-pai-cheng-zui-xiao-de-shu-lcof/solution/mian-shi-ti-45-ba-shu-zu-pai-cheng-zui-xiao-de-s-4/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
func main() {
	fmt.Println(minNumber([]int{12,121}))
	//fmt.Println(minNumber([]int{824,938,1399,5607,6973,5703,9609,4398,8247}))
	//fmt.Println(minNumber([]int{128, 12}))
	fmt.Println(minNumber([]int{3,30,34,5,9}))
}

func minNumber(nums []int) string {
	sArr := make([]string, 0)
	for i := 0; i < len(nums); i++ {
		sArr = append(sArr, strconv.Itoa(nums[i]))
	}

	sort.Slice(sArr, func(i, j int) bool {
		return sArr[i] + sArr[j] < sArr[j] + sArr[i]
	})

	return strings.Join(sArr, "")
}
