package leetcode

import (
	"testing"
)

func TestRob(t *testing.T) {
	inputs := [][]int{
		{1, 2, 3, 1},
		{2, 7, 9, 3, 1},
	}
	outputs := []int{
		4,
		12,
	}
	for k, v := range inputs {
		expected := outputs[k]
		t.Logf("Rob, input: %v, expect: %v", v, expected)

		res := rob(v)
		if res == expected {
			t.Log("Rob is ok")
		} else {
			t.Errorf("Rob is not ok, result:%v", res)
		}
	}
}
