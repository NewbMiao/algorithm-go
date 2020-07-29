package leetcode

import (
	"testing"
)

func TestMaxSubArray(t *testing.T) {
	inputs := [][]int{
		{-2, 1, -3, 4, -1, 2, 1, -5, 4},
	}
	outputs := []int{
		6,
	}
	for k, v := range inputs {
		expected := outputs[k]
		t.Logf("maxSubArray, input: %v, expect: %v", v, expected)
		result := maxSubArray(v)
		if result == expected {
			t.Log("maxSubArray is ok")
		} else {
			t.Errorf("maxSubArray is not ok, result:%v", v)
		}
	}
}
