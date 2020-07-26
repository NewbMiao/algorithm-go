package leetcode

import (
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	inputs := [][]string{
		{"flower", "flow", "flight"},
		{"dog", "racecar", "car"},
	}
	outputs := []string{
		"fl",
		"",
	}
	for k, v := range inputs {
		expected := outputs[k]
		t.Logf("longestCommonPrefix, input: %v, expect: %v", v, expected)
		res := longestCommonPrefix(v)
		if res == expected {
			t.Log("longestCommonPrefix is ok")
		} else {
			t.Errorf("longestCommonPrefix is not ok, result:%v, expect:%v", res, expected)
		}
	}
}
