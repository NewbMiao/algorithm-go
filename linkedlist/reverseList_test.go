package linkedlist

import (
	"fmt"
	"testing"

	"github.com/NewbMiao/algorithm-go/kit/linkedlist"
)

func TestReverseList(t *testing.T) {
	l := linkedlist.NewList()
	l.Push(1)
	l.Push(2)
	l.Push(3)
	l.Push(4)
	l.Push(5)
	want := []int{5, 4, 3, 2, 1}
	t.Log(fmt.Sprintf("ReverseList, input list %v, want: %v", l, want))
	ReverseList(l)
	if fmt.Sprint(l) == fmt.Sprint(want) {
		t.Log("ReverseList is ok")
	} else {
		t.Error("ReverseList is not ok, result:", fmt.Sprintf("%v", l))
	}
}

func TestReverseDList(t *testing.T) {
	l := linkedlist.NewDList()
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l.PushBack(4)
	l.PushBack(5)
	want := []int{5, 4, 3, 2, 1}
	t.Log(fmt.Sprintf("ReverseDList, input list %v, want: %v", l, want))
	ReverseDList(l)
	if fmt.Sprint(l) == fmt.Sprint(want) {
		t.Log("ReverseDList is ok")
	} else {
		t.Error("ReverseDList is not ok, result:", fmt.Sprintf("%v", l))
	}
}
