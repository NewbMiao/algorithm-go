package leetcode

import (
	"fmt"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	ds := [][3][]int{
		{{1, 2, 4}, {1, 3, 4}, {1, 1, 2, 3, 4, 4}},
		{{}, {}, {}},
		{{}, {0}, {0}},
	}
	for _, v := range ds {
		t.Logf("MergeTwoLists, input list num1 %v, num2 %v, want: %v", v[0], v[1], v[2])
		r := mergeTwoLists(generateLinkedListFromArray(v[0]), generateLinkedListFromArray(v[1]))
		tmp := r.convertListNodesToArray()
		if fmt.Sprint(tmp) == fmt.Sprint(v[2]) {
			t.Log("MergeTwoLists is ok")
		} else {
			t.Errorf("MergeTwoLists is not ok, result:%v", tmp)
		}

		r = mergeTwoListsRecursive(generateLinkedListFromArray(v[0]), generateLinkedListFromArray(v[1]))
		tmp = r.convertListNodesToArray()
		if fmt.Sprint(tmp) == fmt.Sprint(v[2]) {
			t.Log("mergeTwoListsRecursive is ok")
		} else {
			t.Errorf("mergeTwoListsRecursive is not ok, result:%v", tmp)
		}
	}
}
