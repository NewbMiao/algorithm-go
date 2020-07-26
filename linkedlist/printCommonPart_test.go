package linkedlist

import (
	"fmt"
	"testing"

	"github.com/NewbMiao/algorithm-go/kit/linkedlist"
)

func TestPrintCommonPart(t *testing.T) {
	h1 := linkedlist.NewList()
	h1.Push(2)
	h1.Push(3)
	h1.Push(5)
	h1.Push(6)

	h2 := linkedlist.NewList()
	h2.Push(1)
	h2.Push(2)
	h2.Push(5)
	h2.Push(7)
	r := printCommonPart(h1, h2)
	want := []int{2, 5}
	t.Log(fmt.Sprintf("PrintCommonPart, input list %v,%v, want: %v", h1, h2, want))
	if fmt.Sprint(r) == fmt.Sprint(want) {
		t.Log("PrintCommonPart is ok")
	} else {
		t.Error("PrintCommonPart is not ok, result:", fmt.Sprintf("%v", r))
	}
}
