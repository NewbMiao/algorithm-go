package stack

import (
	"testing"
	"fmt"
	. "kit/stack"
)

func TestReverse(t *testing.T) {
	st := New()
	want := []int{1, 2, 3, 4, 5}
	res := []int{}
	st.Push(1)
	st.Push(2)
	st.Push(3)
	st.Push(4)
	st.Push(5)

	Reverse(st)

	t.Log(fmt.Sprintf("reverse, want: %v", want))

	if fmt.Sprintf("%v", st.Data) == fmt.Sprintf("%v", want) {
		t.Log("reverse is ok")
	} else {
		t.Error("reverse is not ok, result:", fmt.Sprintf("%v", res))
	}
}
