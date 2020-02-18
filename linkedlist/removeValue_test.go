package linkedlist

import (
	"testing"
	. "github.com/NewbMiao/AlgorithmCodeNote/kit/linkedlist"
	"fmt"
)

func TestRemoveValue(t *testing.T) {
	ds := [][]interface{}{
		{1, 1, 3, 3, 1, 2, 1, 1},
	}
	ws := [][]int{
		{3, 3, 2},
	}
	nums := []int{1}
	for k := range ds {
		l := Arr2LList(ds[k])
		t.Log(fmt.Sprintf("removeValue, input list %v, num %v, want: %v", ds[k], nums[k], ws[k]))
		r1 := removeValue1(l.Head, nums[k])
		tmp := NewList()
		tmp.Head = r1
		if fmt.Sprint(tmp) == fmt.Sprint(ws[k]) {
			t.Log("removeValue1 is ok")
		} else {
			t.Error("removeValue1 is not ok, result:", fmt.Sprintf("%v", tmp))
		}

		l2 := Arr2LList(ds[k])
		r2 := removeValue2(l2.Head, nums[k])
		tmp.Head = r2
		if fmt.Sprint(tmp) == fmt.Sprint(ws[k]) {
			t.Log("removeValue2 is ok")
		} else {
			t.Error("removeValue2 is not ok, result:", fmt.Sprintf("%v", tmp))
		}
	}

}
