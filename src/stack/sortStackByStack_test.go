package stack

import (
	"testing"
	"fmt"
	. "kit/stack"
)

func TestSortStackByStack(t *testing.T) {
	s := New()
	s.Push(3)
	s.Push(1)
	s.Push(6)
	s.Push(2)
	s.Push(5)
	s.Push(4)

	want := []int{6, 5, 4, 3, 2, 1}
	sortStackByStack(s)
	t.Log(fmt.Sprintf("SortStackByStack, want: %v", want))

	if fmt.Sprintf("%v", s.Data) == fmt.Sprintf("%v", want) {
		t.Log("SortStackByStack is ok")
	} else {
		t.Error("SortStackByStack is not ok, result:", fmt.Sprintf("%v", s.Data))
	}
}
