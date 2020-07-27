package linkedlist

import (
	"fmt"
	"testing"

	"github.com/NewbMiao/algorithm-go/kit/linkedlist"
)

func TestRemoveMidNode(t *testing.T) {
	l := linkedlist.NewList()
	l.Push(1)
	l.Push(2)
	l.Push(3)
	l.Push(4)
	l.Push(5)
	l.Push(6)
	want := []int{1, 2, 4, 5, 6}
	t.Log(fmt.Sprintf("RemoveMidNode, input list %v, want: %v", l, want))
	RemoveMidNode(l)
	if fmt.Sprint(l) == fmt.Sprint(want) {
		t.Log("RemoveMidNode is ok")
	} else {
		t.Error("RemoveMidNode is not ok, result:", fmt.Sprintf("%v", l))
	}
}

func TestRemoveByRatio(t *testing.T) {
	l := linkedlist.NewList()
	l.Push(1)
	l.Push(2)
	l.Push(3)
	l.Push(4)
	l.Push(5)
	l.Push(6)
	want := []int{1, 2, 4, 5, 6}
	t.Log(fmt.Sprintf("RemoveByRatio, input list %v, want: %v", l, want))
	RemoveByRatio(l, 2, 5)
	if fmt.Sprint(l) == fmt.Sprint(want) {
		t.Log("RemoveByRatio is ok")
	} else {
		t.Error("RemoveByRatio is not ok, result:", fmt.Sprintf("%v", l))
	}

	want = []int{1, 4, 5, 6}
	t.Log(fmt.Sprintf("RemoveByRatio, input list %v, want: %v", l, want))
	RemoveByRatio(l, 1, 3)
	if fmt.Sprint(l) == fmt.Sprint(want) {
		t.Log("RemoveByRatio is ok")
	} else {
		t.Error("RemoveByRatio is not ok, result:", fmt.Sprintf("%v", l))
	}
}
