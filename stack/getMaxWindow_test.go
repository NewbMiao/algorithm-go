package stack

import (
	"fmt"
	"testing"
)

func TestGetMaxWindow(t *testing.T) {
	arr := []int{4, 3, 5, 4, 3, 3, 6, 7}
	want := []int{5, 5, 5, 4, 6, 7}
	width := 3

	r := getMaxWindow(arr, width)
	t.Log(fmt.Sprintf("GetMaxWindow, input %v, want: %v", arr, want))

	if fmt.Sprintf("%v", r) == fmt.Sprintf("%v", want) {
		t.Log("GetMaxWindow is ok")
	} else {
		t.Error("GetMaxWindow is not ok, result:", fmt.Sprintf("%v", r))
	}
}
