package leetcode

import (
	"testing"
)

func TestMinPathSum(t *testing.T) {
	inputs := [][][]int{
		{
			{1, 3, 1},
			{1, 5, 1},
			{4, 2, 1},
		},
	}
	outputs := []int{
		7,
	}
	for k, v := range inputs {
		expected := outputs[k]
		t.Logf("minPathSum, input: %v, expect: %v", v, expected)
		n := minPathSum(v)
		if n == expected {
			t.Log("minPathSum is ok")
		} else {
			t.Errorf("minPathSum is not ok, result:%v", n)
		}
	}
}
