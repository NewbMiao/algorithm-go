package subsequence

import (
	"testing"
	"fmt"
)

func TestGetLIS(t *testing.T) {
	arr := []int{2, 1, 5, 3, 6, 4, 8, 9, 7}
	want := []int{1, 3, 4, 8, 9}
	t.Log(fmt.Sprintf("GetLIS, input %v, want: %v", arr, want))

	r := GetLIS1(arr)
	if fmt.Sprint(r) == fmt.Sprint(want) {
		t.Log("GetLIS1 is ok")
	} else {
		t.Error("GetLIS1 is not ok, result:", fmt.Sprintf("%v", r))
	}

	r = GetLIS2(arr)
	if fmt.Sprint(r) == fmt.Sprint(want) {
		t.Log("GetLIS2 is ok")
	} else {
		t.Error("GetLIS2 is not ok, result:", fmt.Sprintf("%v", r))
	}
}
