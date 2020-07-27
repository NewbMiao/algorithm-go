package array

import (
	"fmt"
	"testing"
)

func TestLeftUnique(t *testing.T) {
	d := []int{1, 2, 2, 2, 3, 3, 4, 5, 6, 6, 7, 7, 8, 8, 8, 9}
	want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 6, 2, 7, 2, 8, 8, 3}
	t.Log(fmt.Sprintf("partitionArray_leftUnique, input list %v, want: %v", d, want))
	leftUnique(d)
	if fmt.Sprint(want) == fmt.Sprint(d) {
		t.Log("leftUnique is ok")
	} else {
		t.Error("leftUnique is not ok, result:", d)
	}
}

func TestSort(t *testing.T) {
	d := []int{2, 1, 2, 0, 1, 1, 2, 2, 0}
	want := []int{0, 0, 1, 1, 1, 2, 2, 2, 2}
	t.Log(fmt.Sprintf("partitionArray_sort, input list %v, want: %v", d, want))
	pSort(d)
	if fmt.Sprint(want) == fmt.Sprint(d) {
		t.Log("sort is ok")
	} else {
		t.Error("sort is not ok, result:", d)
	}
}
