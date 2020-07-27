package leetcode

import (
	"fmt"
	"testing"
)

func TestPlusOne(t *testing.T) {
	inputs := [][]int{
		{1, 2, 3},
		{4, 3, 2, 1},
		{9, 9, 9},
	}
	outputs := [][]int{
		{1, 2, 4},
		{4, 3, 2, 2},
		{1, 0, 0, 0},
	}
	for k, v := range inputs {
		expected := outputs[k]
		t.Logf("plusOne, input: %v, expect: %v", v, expected)
		res := plusOne(v)
		if fmt.Sprint(res) == fmt.Sprint(expected) {
			t.Log("plusOne is ok")
		} else {
			t.Errorf("plusOne is not ok, result:%v, expect:%v", res, expected)
		}
	}
}
