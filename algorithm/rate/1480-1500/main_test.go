package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)


func Test_numSplits(t *testing.T) {
	m := map[string]int{
		"aacaba":2,
		"abcd":1,
		"aaaaa":4,
	}

	for k, v := range m {
		assert.Equal(t, numSplits(k), v, "error,不相等")
	}
}

func Test_rearrangeArray(t *testing.T)  {
	arr := [][]int{
		{1,2,3,4,5},
		{6,2,0,9,7},
	}

	for i := 0; i < len(arr); i++ {
		t.Log(rearrangeArray(arr[i]))
	}
}



