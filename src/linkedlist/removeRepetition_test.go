package linkedlist

import (
	"fmt"
	"testing"
	. "kit/linkedlist"
)

func TestRemoveRepetition(t *testing.T) {
	ds := [][]interface{}{
		{1, 2, 3, 3, 4, 4, 2, 1, 1},
		{1, 1, 3, 3, 4, 4, 2, 1, 1},
	}
	ws := [][]int{
		{1, 2, 3, 4},
		{1, 3, 4, 2},
	}
	for k := range ds {
		l := Arr2LList(ds[k])
		t.Log(fmt.Sprintf("RemoveRepetition1, input list %v, want: %v", ds[k], ws[k]))
		removeRepetition1(l.Head)
		if fmt.Sprint(l) == fmt.Sprint(ws[k]) {
			t.Log("RemoveRepetition is ok")
		} else {
			t.Error("RemoveRepetition1 is not ok, result:", fmt.Sprintf("%v", l))
		}

		l2 := Arr2LList(ds[k])
		removeRepetition2(l2.Head)
		if fmt.Sprint(l2) == fmt.Sprint(ws[k]) {
			t.Log("RemoveRepetition2 is ok")
		} else {
			t.Error("RemoveRepetition2 is not ok, result:", fmt.Sprintf("%v", l2))
		}
	}

}
