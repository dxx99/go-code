package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println(largestWordCount([]string{"Hello userTwooo","Hi userThree","Wonderful day Alice","Nice day userThree"}, []string{"Alice","userTwo","userThree","Alice"}))
	fmt.Println(largestWordCount([]string{"How is leetcode for everyone","Leetcode is useful for practice"}, []string{"Bob","Charlie"}))
}



func largestWordCount(messages []string, senders []string) string {
	maxName := ""
	maxCount := math.MinInt
	hash := make(map[string]int)
	for i := 0; i < len(senders); i++ {
		if v, ok := hash[senders[i]]; ok {
			hash[senders[i]] = v + len(strings.Split(messages[i], " "))
		} else {
			hash[senders[i]] = len(strings.Split(messages[i], " "))
		}
		if hash[senders[i]] > maxCount || (hash[senders[i]] == maxCount && senders[i] > maxName) {
			maxName = senders[i]
			maxCount = hash[senders[i]]
		}
	}

	return maxName
}