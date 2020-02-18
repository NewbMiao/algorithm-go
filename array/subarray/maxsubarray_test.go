package subarray

import (
	"testing"
	"fmt"
)

func TestMaxSubArraySum(t *testing.T) {
	data := []int{1, 2, 5, 1, 3}
	sum := 6
	want := []int{5, 1}
	t.Log(fmt.Sprintf("MaxSubArraySum %v, want subArray which sum is 6: %v", data, want))
	r := MaxSubArraySum(data, sum)
	if fmt.Sprintf("%v", r) == fmt.Sprintf("%v", want) {
		t.Log("MaxSubArraySum is ok")
	} else {
		t.Error("MaxSubArraySum is not ok, result:", r)
	}
}
