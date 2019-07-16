package stack

import (
	"testing"
	"fmt"
)

func TestMaxRecSize(t *testing.T) {
	d := [][]int{
		{1, 0, 1, 1},
		{1, 1, 1, 1},
		{1, 1, 1, 0},
	}
	want := 6
	r := maxRecSize(d)

	t.Log(fmt.Sprintf("maxRecSize, input: %v, want: %v", d, want))

	if r == want {
		t.Log("maxRecSize is ok")
	} else {
		t.Error("maxRecSize is not ok, result:", fmt.Sprintf("%v", r))
	}
}
