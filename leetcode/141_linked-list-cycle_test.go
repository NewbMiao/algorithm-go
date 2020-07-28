package leetcode

import (
	"testing"
)

func TestHasCycle(t *testing.T) {
	type input struct {
		nums []int
		pos  int
	}
	inputs := []input{
		{[]int{3, 2, 0, 4}, 1},
		{[]int{1, 2}, 0},
		{[]int{1, 2}, -1},
	}
	outputs := []bool{
		true,
		true,
		false,
	}
	for k, v := range inputs {
		t.Logf("hasCycle, input list %v , want: %v", v, outputs[k])
		r := hasCycle(generateCycleLinkedList(v.nums, v.pos))
		if r == outputs[k] {
			t.Log("hasCycle is ok")
		} else {
			t.Errorf("hasCycle is not ok, result:%v", r)
		}
	}
}
