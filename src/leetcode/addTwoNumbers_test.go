package leetcode

import (
	"testing"
	"fmt"
)

func TestAddTwoNumbers(t *testing.T) {
	ds := [][3][]int{
		{{2, 4, 3}, {5, 6, 4}, {7, 0, 8}},
		{{5}, {5}, {0, 1}},
		{{1,8},{0},{1,8}},
	}
	for _, v := range ds {
		fmt.Println(v)
		t.Log(fmt.Sprintf("AddTwoNumbers, input list num1 %v, num2 %v, want: %v", v[0], v[1], v[2]))
		fmt.Println(genListNode(v[0]), cvtListNode2arr(genListNode(v[0])))
		r := addTwoNumbers(genListNode(v[0]), genListNode(v[1]))
		tmp := cvtListNode2arr(r)
		if fmt.Sprint(tmp) == fmt.Sprint(v[2]) {
			t.Log("AddTwoNumbers is ok")
		} else {
			t.Error("AddTwoNumbers is not ok, result:", fmt.Sprintf("%v", tmp))
		}

	}

}

func genListNode(arr []int) *ListNode {
	l := &ListNode{}
	cur := l
	for i := 0; i < len(arr); i++ {
		cur.Val = arr[i]
		if i == len(arr)-1 {
			break
		}
		cur.Next = &ListNode{}
		cur = cur.Next
	}
	return l
}

func cvtListNode2arr(l *ListNode) (res []int) {
	cur := l
	for cur != nil {
		res = append(res, cur.Val)
		cur = cur.Next
	}
	return
}
