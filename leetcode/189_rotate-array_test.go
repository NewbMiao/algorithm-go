package leetcode

import (
	"fmt"
	"testing"
)

func TestRotate(t *testing.T) {
	type input struct {
		nums   []int
		target int
	}
	inputs := []input{
		{[]int{1, 2, 3, 4, 5, 6, 7}, 3},
		{[]int{-1, -100, 3, 99}, 2},
	}
	outputs := [][]int{
		{5, 6, 7, 1, 2, 3, 4},
		{3, 99, -1, -100},
	}
	for k, v := range inputs {
		expected := outputs[k]
		t.Logf("rotate, input: %v, expect: %v", v, expected)

		inputCopy := make([]int, len(v.nums))
		copy(inputCopy, v.nums)

		rotate(v.nums, v.target)
		if fmt.Sprint(v.nums) == fmt.Sprint(expected) {
			t.Log("rotate is ok")
		} else {
			t.Errorf("rotate is not ok, result:%v, expect:%v", v.nums, expected)
		}
		rotateAppend(inputCopy, v.target)
		if fmt.Sprint(inputCopy) == fmt.Sprint(expected) {
			t.Log("rotate_append is ok")
		} else {
			t.Errorf("rotate_append is not ok, result:%v, expect:%v", inputCopy, expected)
		}
	}
}
