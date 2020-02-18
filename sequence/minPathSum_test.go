package sequence

import (
	"testing"
	"fmt"
)

func TestMinPathSum(t *testing.T) {
	m := [][]int{
		{1, 3, 5, 9},
		{8, 1, 3, 4},
		{5, 0, 6, 1},
		{8, 8, 4, 0},
	}
	mStr := "\n"
	for _, v := range m {
		mStr += fmt.Sprint(v) + "\n"
	}
	want := 12
	var r int
	t.Log(fmt.Sprintf("input %v minPathSum want: %v", mStr, want))

	r = MinPathSum1(m)
	if r == want {
		t.Log("minPathSum1 is ok")
	} else {
		t.Error("minPathSum1 is not ok, result:", fmt.Sprintf("%v", r))
	}

	r = MinPathSum2(m)
	if r == want {
		t.Log("minPathSum2 is ok")
	} else {
		t.Error("minPathSum2 is not ok, result:", fmt.Sprintf("%v", r))
	}
}
