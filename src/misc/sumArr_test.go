package misc

import (
	"testing"
	"sort"
)

func TestSumArr(t *testing.T) {
	arr := []int{1, 10, 4, 22, 5, 6, 7, 8}
	sort.Ints(arr)
	sum := 20
	res := []int{}
	index := len(arr) - 1
	findSumArr(arr, res, sum, index)
	/*
	output:
		10,6,4
		10,5,4,1
		8,7,5
		8,7,4,1
		8,6,5,1
	*/
}
