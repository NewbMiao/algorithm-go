package leetcode

import (
	"testing"
	"fmt"
)

func TestMergeKLists(t *testing.T) {
	ds := [][][]int{
		{{1, 4, 5}, {1, 3, 4}, {2, 6},{7,8}},
		//{{}, {}, {}},
		//{{}, {0}, {0}},
	}
	ws := [][]int{
		{1, 1, 2, 3, 4, 4, 5, 6,7,8},
		{},
		{0,0},
	}
	for k, v := range ds {
		t.Log(fmt.Sprintf("MergeKLists, input lists  %v, want: %v", v, ws[k]))

		r := mergeKLists(fmtLists(v))
		tmp := cvtListNode2arr(r)
		if fmt.Sprint(tmp) == fmt.Sprint(ws[k]) {
			t.Log("MergeKLists is ok")
		} else {
			t.Error("MergeKLists is not ok, result:", fmt.Sprintf("%v", tmp))
		}

	}
}

func fmtLists(list [][]int) (res []*ListNode) {
	for _, v := range list {
		res = append(res, genListNode(v))
	}
	return
}
