package leetcode

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	type input struct {
		nums   []int
		target int
	}
	inputs := []input{
		{[]int{2, 7, 11, 15}, 9},
		{[]int{2, 7, 11, 15}, 22},
	}
	outputs := [][]int{
		{0, 1},
		{1, 3},
	}
	for k, v := range inputs {
		expected := outputs[k]
		t.Logf("twoSum, input: %v, expect: %v", v, expected)
		res := twoSum(v.nums, v.target)
		if fmt.Sprint(res) == fmt.Sprint(expected) {
			t.Log("twoSum is ok")
		} else {
			t.Errorf("twoSum is not ok, result:%v, expect:%v", res, expected)
		}
	}
}
