package leetcode

import (
	"fmt"
	"testing"
)

func TestMaxArea(t *testing.T) {
	d := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	want := 49
	t.Log(fmt.Sprintf("maxArea, input arr %v, want: %v", d, want))
	r := maxArea(d)
	if r == want {
		t.Log("maxArea is ok")
	} else {
		t.Error("maxArea is not ok, result:", fmt.Sprintf("%v", r))
	}
}
