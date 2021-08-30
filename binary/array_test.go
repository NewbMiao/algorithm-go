package binary

import (
	"fmt"
	"testing"
)

func TestGetNoRepeatNum(t *testing.T) {
	arr := []int{
		1, 2, 3, 4, 5, 1, 2, 3, 4,
	}
	want := 5
	t.Log(fmt.Sprintf("input x=%v, no repeat num is : %d", arr, want))

	if getNoRepeatNum(arr) == want {
		t.Log("getNoRepeatNum is ok")
	} else {
		t.Error("getNoRepeatNum is not ok")
	}
}
