package leetcode

import (
	"testing"
)

func TestFirstUniqChar(t *testing.T) {
	inputs := []string{
		"leetcode",
		"loveleetcode",
	}
	outputs := []int{
		0,
		2,
	}
	for k, v := range inputs {
		expected := outputs[k]
		t.Logf("FirstUniqChar, input: %v, expect: %v", v, expected)

		res := firstUniqChar(v)
		if res == expected {
			t.Log("FirstUniqChar is ok")
		} else {
			t.Errorf("FirstUniqChar is not ok, result:%v", res)
		}
	}
}
