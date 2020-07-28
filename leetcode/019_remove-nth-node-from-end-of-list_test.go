package leetcode

import (
	"fmt"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	ds := [][]int{
		{1, 2, 3, 4, 5},
		{1},
	}
	ns := []int{
		2,
		1,
	}
	ws := [][]int{
		{1, 2, 3, 5},
		{},
	}
	for k, v := range ds {
		t.Logf("RemoveNthFromEnd, input list %v, del last %dth node , want: %v", v, ns[k], ws[k])
		r := removeNthFromEnd(generateLinkedListFromArray(v), ns[k])
		tmp := r.convertListNodesToArray()
		if fmt.Sprint(tmp) == fmt.Sprint(ws[k]) {
			t.Log("RemoveNthFromEnd is ok")
		} else {
			t.Errorf("RemoveNthFromEnd is not ok, result:%v", tmp)
		}

		r = removeNthFromEnd2(generateLinkedListFromArray(v), ns[k])
		tmp = r.convertListNodesToArray()
		if fmt.Sprint(tmp) == fmt.Sprint(ws[k]) {
			t.Log("RemoveNthFromEnd is ok")
		} else {
			t.Errorf("RemoveNthFromEnd is not ok, result:%v", tmp)
		}
	}
}
