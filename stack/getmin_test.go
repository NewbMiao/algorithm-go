package stack

import (
	"fmt"
	"testing"
)

func TestMyStack_GetMin(t *testing.T) {
	ms := NewMyStack()
	want := []int{3, 3, 1, 3}
	res := []int{}
	ms.Push(3)
	res = append(res, ms.GetMin())
	ms.Push(4)
	res = append(res, ms.GetMin())

	ms.Push(1)
	res = append(res, ms.GetMin())

	ms.Pop()
	res = append(res, ms.GetMin())
	t.Log(fmt.Sprintf("getMin, want: %v", want))

	if fmt.Sprintf("%v", res) == fmt.Sprintf("%v", want) {
		t.Log("getMin is ok")
	} else {
		t.Error("getMin is not ok, result:", fmt.Sprintf("%v", res))
	}
}
