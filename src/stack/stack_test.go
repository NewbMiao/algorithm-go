package stack

import (
	"testing"
	"fmt"
)

func TestStack(t *testing.T) {
	want := []int{
		1, 2, 3, 4, 5,
	}
	s := &S{}
	s.Push(5)
	s.Push(4)
	s.Push(3)
	s.Push(2)
	s.Push(1)
	fmt.Println(s.Data)
	t.Log(fmt.Sprintf("input stack=%v, pop want: %v", s.Data, want))
	var r []int

	for !s.IsEmpty() {
		tmp, _ := s.Pop()
		r = append(r, tmp.(int))
	}
	if fmt.Sprintf("%v", r) == fmt.Sprintf("%v", want) {
		t.Log("stack is ok")
	} else {
		t.Error("stack is not ok, result:", fmt.Sprintf("%v", r))
	}
}
