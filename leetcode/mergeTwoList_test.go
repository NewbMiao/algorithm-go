package leetcode

import (
	"testing"
	"fmt"
)

func TestMergeTwoLists(t *testing.T) {
	ds := [][3][]int{
		{{1, 2, 4,}, {1, 3, 4}, {1, 1, 2, 3, 4, 4}},
		{{},{},{}},
		{{},{0},{0}},
	}
	for _, v := range ds {
		t.Log(fmt.Sprintf("MergeTwoLists, input list num1 %v, num2 %v, want: %v", v[0], v[1], v[2]))
		fmt.Println(genListNode(v[0]), cvtListNode2arr(genListNode(v[0])))
		r := mergeTwoLists(genListNode(v[0]), genListNode(v[1]))
		tmp := cvtListNode2arr(r)
		if fmt.Sprint(tmp) == fmt.Sprint(v[2]) {
			t.Log("MergeTwoLists is ok")
		} else {
			t.Error("MergeTwoLists is not ok, result:", fmt.Sprintf("%v", tmp))
		}

	}
}
