package leetcode

import (
	"fmt"
	"sort"
	"testing"
)

func TestIntersect(t *testing.T) {
	inputs := [][2][]int{
		{{1, 2, 2, 1}, {2, 2}},
		{{4, 9, 5}, {9, 4, 9, 8, 4}},
	}
	outputs := [][]int{
		{2, 2},
		{4, 9},
	}
	for k, v := range inputs {
		expected := outputs[k]
		t.Logf("intersect, input: %v, expect: %v", v, expected)
		res := intersect(v[0], v[1])
		sort.Ints(res)
		if fmt.Sprint(res) == fmt.Sprint(expected) {
			t.Log("intersect is ok")
		} else {
			t.Errorf("intersect is not ok, result:%v, expect:%v", res, expected)
		}

		res = intersectDoublepointer(v[0], v[1])
		sort.Ints(res)
		if fmt.Sprint(res) == fmt.Sprint(expected) {
			t.Log("intersect_doublePointer is ok")
		} else {
			t.Errorf("intersect_doublePointer is not ok, result:%v, expect:%v", res, expected)
		}
	}
}
