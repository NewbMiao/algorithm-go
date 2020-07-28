package leetcode

import (
	"fmt"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	inputs := [][2][]int{
		{{2, 4, 3}, {5, 6, 4}},
	}
	outputs := [][]int{
		{7, 0, 8},
	}
	for k, v := range inputs {
		expected := outputs[k]
		t.Logf("addTwoNumbers, input: %v, expect: %v", v, expected)
		res := addTwoNumbers(generateLinkedListFromArray(v[0]), generateLinkedListFromArray(v[1]))
		if fmt.Sprint(res.convertListNodesToArray()) == fmt.Sprint(expected) {
			t.Log("addTwoNumbers is ok")
		} else {
			t.Errorf("addTwoNumbers is not ok, result:%v, expect:%v", res, expected)
		}
	}
}
