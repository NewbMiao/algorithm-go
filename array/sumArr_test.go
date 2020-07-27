package array

import (
	"fmt"
	"testing"
)

func TestSumArr(t *testing.T) {
	arr := []int{1, 10, 4, 22, 5, 6}
	sum := 20
	r := FindSumArr(arr, sum)
	t.Log(fmt.Sprintf("SumArr, input arr :%v, want subArray which sum is %d : %v", arr, sum, r))

	/*
		output:
			10,6,4
			10,5,4,1
			8,7,5
			8,7,4,1
			8,6,5,1
	*/
}
