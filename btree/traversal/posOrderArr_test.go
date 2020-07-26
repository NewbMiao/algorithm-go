package traversal

import (
	"fmt"
	"testing"

	"github.com/NewbMiao/algorithm-go/kit/btree"
)

func TestIsPosOrderArr(t *testing.T) {
	data := []int{2, 1, 3, 6, 5, 7, 4}
	want := true
	t.Log(fmt.Sprintf("input pos order arr %v, IsPosOrderArr want: %v", data, want))
	r := IsPosOrderArr(data)
	if r == want {
		t.Log("IsPosOrderArr is ok")
	} else {
		t.Error("IsPosOrderArr is not ok, result:", r)
	}
	btree.PrintTree(PosOrderArrToBST(data))
}
