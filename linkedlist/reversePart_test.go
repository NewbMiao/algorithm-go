package linkedlist

import (
	"fmt"
	"testing"

	"github.com/NewbMiao/algorithm-go/kit/linkedlist"
)

func TestReversePart(t *testing.T) {
	l := linkedlist.NewList()
	l.Push(1)
	l.Push(2)
	l.Push(3)
	l.Push(4)
	l.Push(5)
	want := []int{3, 2, 1, 4, 5}
	t.Log(fmt.Sprintf("ReversePart, input list %v, want: %v", l, want))
	ReversePart(l, 1, 3)
	if fmt.Sprint(l) == fmt.Sprint(want) {
		t.Log("ReversePart is ok")
	} else {
		t.Error("ReversePart is not ok, result:", fmt.Sprintf("%v", l))
	}
}
