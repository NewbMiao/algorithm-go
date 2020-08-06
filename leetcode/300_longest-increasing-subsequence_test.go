package leetcode

import (
	"testing"
)

func TestLengthOfLIS(t *testing.T) {
	inputs := [][]int{
		{10, 9, 2, 5, 3, 7, 101, 18},
		{1, 3, 6, 7, 9, 4, 10, 5, 6},
	}
	outputs := []int{
		4,
		6,
	}
	for k, v := range inputs {
		expected := outputs[k]
		t.Logf("lengthOfLIS, input: %v, expect: %v", v, expected)

		res := lengthOfLIS(v)
		if res == expected {
			t.Log("lengthOfLIS is ok")
		} else {
			t.Errorf("lengthOfLIS is not ok, result:%v", res)
		}
		res = lengthOfLISGreedy(v)
		if res == expected {
			t.Log("lengthOfLISGreedy is ok")
		} else {
			t.Errorf("lengthOfLISGreedy is not ok, result:%v", res)
		}
	}
}
