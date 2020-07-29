package leetcode

import (
	"testing"
)

func TestClimbStairs(t *testing.T) {
	inputs := []int{
		2,
		3,
	}
	outputs := []int{
		2,
		3,
	}
	for k, v := range inputs {
		expected := outputs[k]
		t.Logf("climbStairs, input: %v, expect: %v", v, expected)
		n := climbStairs(v)
		if n == expected {
			t.Log("climbStairs is ok")
		} else {
			t.Errorf("climbStairs is not ok, result:%v", v)
		}
	}
}
