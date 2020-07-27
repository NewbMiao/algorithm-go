package linkedlist

import (
	"fmt"
	"testing"

	"github.com/NewbMiao/algorithm-go/kit/linkedlist"
)

func TestListPartition(t *testing.T) {
	d := []interface{}{7, 9, 1, 8, 5, 2, 5}
	pivot := 4
	want := []int{2, 1, 8, 5, 9, 5, 7}
	l := linkedlist.Arr2LList(d)
	t.Log(fmt.Sprintf("listPartition1, input list %v, pivot %d, want: %v", l, pivot, want))
	listPartition1(l, pivot)
	if fmt.Sprint(want) == fmt.Sprint(l) {
		t.Log("listPartition1 is ok")
	} else {
		t.Error("listPartition1 is not ok, result:", l)
	}

	l = linkedlist.Arr2LList(d)
	pivot = 5
	want = []int{1, 2, 5, 5, 7, 9, 8}
	t.Log(fmt.Sprintf("listPartition2 with order, input list %v, pivot %d, want: %v", l, pivot, want))
	listPartition2(l, pivot)
	if fmt.Sprint(want) == fmt.Sprint(l) {
		t.Log("listPartition2 is ok")
	} else {
		t.Error("listPartition2 is not ok, result:", l)
	}
}
