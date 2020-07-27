package leetcode

import (
	"fmt"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	type input struct {
		nums   []int
		target int
	}
	inputs := []input{
		{[]int{3, 2, 2, 3}, 3},
		{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2},
	}
	outputs := [][]int{
		{2, 2},
		{0, 1, 3, 0, 4},
	}
	for k, v := range inputs {
		expected := outputs[k]
		t.Logf("removeElement, input: %v, expect: %v", v, expected)
		length := removeElement(v.nums, v.target)
		if fmt.Sprint(v.nums[:length]) == fmt.Sprint(expected) {
			t.Log("removeElement is ok")
		} else {
			t.Errorf("removeElement is not ok, result:%v, expect:%v", v.nums[:length], expected)
		}
	}
}
