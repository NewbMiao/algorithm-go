package linkedlist

import (
	"fmt"
	"testing"

	"github.com/NewbMiao/algorithm-go/kit/linkedlist"
)

func TestReverseKNodes(t *testing.T) {
	ds := [][]interface{}{
		{},
		{1},
		{1, 2},
		{1, 2},
		{1, 2, 3, 4},
		{1, 2, 3, 4, 5, 6, 7, 8},
	}
	ks := []int{
		3,
		3,
		2,
		3,
		2,
		3,
	}
	ws := [][]int{
		{},
		{1},
		{2, 1},
		{1, 2},
		{2, 1, 4, 3},
		{3, 2, 1, 6, 5, 4, 7, 8},
	}
	for k := range ds {
		l := linkedlist.Arr2LList(ds[k])
		var head *linkedlist.LNode
		if l == nil {
			head = nil
		} else {
			head = l.Head
		}
		t.Log(fmt.Sprintf("reverseKNodes, input list %v, k %v, want: %v", l, ks[k], ws[k]))
		r := reverseKNodes1(head, ks[k])
		if fmt.Sprint(linkedlist.CvtLNode2arr(r)) == fmt.Sprint(ws[k]) {
			t.Log("reverseKNodes1 is ok")
		} else {
			t.Error("reverseKNodes1 is not ok, result:", fmt.Sprintf("%v", r))
		}

		l = linkedlist.Arr2LList(ds[k])
		if l == nil {
			head = nil
		} else {
			head = l.Head
		}
		r = reverseKNodes2(head, ks[k])
		if fmt.Sprint(linkedlist.CvtLNode2arr(r)) == fmt.Sprint(ws[k]) {
			t.Log("reverseKNodes2 is ok")
		} else {
			t.Error("reverseKNodes2 is not ok, result:", fmt.Sprintf("%v", r))
		}

		l = linkedlist.Arr2LList(ds[k])
		if l == nil {
			head = nil
		} else {
			head = l.Head
		}
		r = reverseKNodes3(head, ks[k])
		if fmt.Sprint(linkedlist.CvtLNode2arr(r)) == fmt.Sprint(ws[k]) {
			t.Log("reverseKNodes3 is ok")
		} else {
			t.Error("reverseKNodes3 is not ok, result:", fmt.Sprintf("%v", linkedlist.CvtLNode2arr(r)))
		}
	}
}
