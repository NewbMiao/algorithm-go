package leetcode

import (
	"testing"
)

func TestReverseString(t *testing.T) {
	inputs := [][]byte{
		{'h', 'e', 'l', 'l', 'o'},
		{'H', 'a', 'n', 'n', 'a', 'h'},
	}
	outputs := [][]byte{
		{'o', 'l', 'l', 'e', 'h'},
		{'h', 'a', 'n', 'n', 'a', 'H'},
	}
	for k, v := range inputs {
		expected := outputs[k]
		t.Logf("ReverseString, input: %v, expect: %v", v, expected)

		reverseString(v)
		if string(v) == string(expected) {
			t.Log("ReverseString is ok")
		} else {
			t.Errorf("ReverseString is not ok, result:%v", v)
		}
	}
}
