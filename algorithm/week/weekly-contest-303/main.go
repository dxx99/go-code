package main

import (
	"fmt"
	"strconv"
)

func main() {
	obj := Constructor([]string{"kimchi", "miso", "sushi", "moussaka", "ramen", "bulgogi"}, []string{"korean", "japanese", "japanese", "greek", "japanese", "korean"}, []int{9, 12, 8, 15, 14, 7})
	obj.HighestRated("korean")
	obj.HighestRated("japanese")
	obj.ChangeRating("sushi", 16)
	obj.HighestRated("japanese")
	obj.ChangeRating("ramen", 16)
	obj.HighestRated("japanese")

}

//1.
func repeatedCharacter(s string) byte {
	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		if _, ok := m[s[i]]; ok {
			return s[i]
		}else {
			m[s[i]]++
		}
	}
	return 0
}

//2.
func equalPairs(grid [][]int) int {
	ans := 0
	m := make(map[string]int)
	for i := 0; i < len(grid); i++ {
		key := ""
		for j := 0; j < len(grid[i]); j++ {
			key += "-"+strconv.Itoa(grid[i][j])
		}
		m[key]++
	}
	fmt.Println(m)

	k := 0
	for k < len(grid) {
		tmp := ""
		for j := 0; j < len(grid); j++ {
			tmp += "-"+ strconv.Itoa(grid[j][k])
		}
		fmt.Println(tmp)
		if v, ok := m[tmp]; ok {
			ans += v
		}
		k++
	}
	return ans
}

//3.
type FoodRatings struct {
	fMap map[string]int			// 食物分数
	fc 	 map[string]string		// 食物口味
	cMap map[string][]string	// 口味食物
	cMax map[string]string		// 得分最大的字典树最小的食物
}


func Constructor(foods []string, cuisines []string, ratings []int) FoodRatings {
	fm, fc := make(map[string]int), make(map[string]string)
	for i := 0; i < len(foods); i++ {
		fm[foods[i]] = ratings[i]
		fc[foods[i]] = cuisines[i]
	}

	cmax := make(map[string]string)
	cm := make(map[string][]string)
	for j := 0; j < len(cuisines); j++ {
		if v, ok := cm[cuisines[j]]; ok {
			if (fm[foods[j]] > fm[cmax[cuisines[j]]]) || (fm[foods[j]] == fm[cmax[cuisines[j]]] && foods[j] < cmax[cuisines[j]]) {
				cmax[cuisines[j]] = foods[j]
			}
			cm[cuisines[j]] = append(v, foods[j])
		}else {
			cm[cuisines[j]] = []string{foods[j]}
			cmax[cuisines[j]] = foods[j]
		}
	}

	return FoodRatings{
		fMap: fm,
		fc: fc,
		cMap: cm,
		cMax: cmax,
	}
}


func (f *FoodRatings) ChangeRating(food string, newRating int)  {
	su := f.fc[food]	// 食物的口味
	f.fMap[food] = newRating
	fmt.Println(food, newRating, su, f.cMax[su])
	fmt.Println(f.cMax[su], f.fMap[f.cMax[su]])

	if (f.fMap[f.cMax[su]] < newRating) || (f.fMap[f.cMax[su]] == newRating && f.cMax[su] > food) {		// 如果相等则要重新再找一个最大值

		cs, ok := f.cMap[su]
		if !ok || len(cs) == 0 {
			f.cMax[su] = ""
			return
		}

		maxStr := cs[0]
		for i := 1; i < len(cs); i++ {
			if f.fMap[cs[i]] > f.fMap[maxStr] {
				maxStr = cs[i]
			}else if f.fMap[cs[i]] == f.fMap[maxStr] {
				if cs[i] < maxStr {
					maxStr = cs[i]
				}
			}
		}
		f.cMax[su] = maxStr
	}
}


func (f *FoodRatings) HighestRated(cuisine string) string {
	fmt.Println(cuisine, f.cMax)
	cs, ok := f.cMax[cuisine]
	if !ok {
		return ""
	}
	return cs
}

//4.

