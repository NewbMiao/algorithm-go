package leetcode

import (
	"testing"
	"fmt"
)

func TestTrap(t *testing.T) {
	ds := [][]int{
		{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
	}

	ws := []int{
		6,
	}

	for k, v := range ds {
		t.Log(fmt.Sprintf("trap, input arr %v, want: %v", v, ws[k]))
		r := trap(v)
		if r == ws[k] {
			t.Log("trap is ok")
		} else {
			t.Error("trap is not ok, result:", r)
		}

		r = trap1(v)
		if r == ws[k] {
			t.Log("trap1 is ok")
		} else {
			t.Error("trap1 is not ok, result:", r)
		}
	}
}
