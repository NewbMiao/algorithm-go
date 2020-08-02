package leetcode

import (
	"testing"
)

func TestMinimumTotal(t *testing.T) {
	inputs := [][][]int{
		{
			{2},
			{3, 4},
			{6, 5, 7},
			{4, 1, 8, 3},
		},
		{{-10}},
		{{-1}, {-2, -3}},
		{{-1}, {2, 3}, {1, -1, -3}},
	}
	outputs := []int{
		11,
		-10,
		-4,
		-1,
	}
	for k, v := range inputs {
		expected := outputs[k]
		t.Logf("minimumTotal, input: %v, expect: %v", v, expected)
		n := minimumTotal(v)
		if n == expected {
			t.Log("minimumTotal is ok")
		} else {
			t.Errorf("minimumTotal is not ok, result:%v", n)
		}
	}
}
