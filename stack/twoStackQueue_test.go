package stack

import (
	"fmt"
	"testing"
)

func TestTwoStackQueue(t *testing.T) {
	tq := NewTwoStackQueue()
	want := []int{1, 1, 2, 2, 3, 3}
	res := []int{}
	tq.Add(1)
	tq.Add(2)
	tq.Add(3)
	res = append(res, tq.Peek())
	res = append(res, tq.Poll())
	res = append(res, tq.Peek())
	res = append(res, tq.Poll())
	res = append(res, tq.Peek())
	res = append(res, tq.Poll())
	t.Log(fmt.Sprintf("twoStackQueue, want: %v", want))

	if fmt.Sprintf("%v", res) == fmt.Sprintf("%v", want) {
		t.Log("twoStackQueue is ok")
	} else {
		t.Error("twoStackQueue is not ok, result:", fmt.Sprintf("%v", res))
	}
}
