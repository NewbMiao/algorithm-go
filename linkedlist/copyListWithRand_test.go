package linkedlist

import (
	"testing"
	"fmt"
)

func TestCopyListWithRand(t *testing.T) {
	head := &randNode{Value: 1}
	head.Next = &randNode{Value: 2}
	head.Next.Next = &randNode{Value: 3}
	head.Next.Next.Next = &randNode{Value: 4}
	head.Next.Next.Next.Next = &randNode{Value: 5}
	head.Next.Next.Next.Next.Next = &randNode{Value: 6}

	head.Rand = head.Next.Next.Next.Next.Next                // 1 -> 6
	head.Next.Rand = head.Next.Next.Next.Next.Next           // 2 -> 6
	head.Next.Next.Rand = head.Next.Next.Next.Next           // 3 -> 5
	head.Next.Next.Next.Rand = head.Next.Next                // 4 -> 3
	head.Next.Next.Next.Next.Rand = nil                      // 5 -> null
	head.Next.Next.Next.Next.Next.Rand = head.Next.Next.Next // 6 -> 4

	want := [][3]int{
		{1, 2, 6}, {2, 3, 6}, {3, 4, 5}, {4, 5, 3}, {5, 6, 0}, {6, 0, 4},
	}

	t.Log(fmt.Sprintf("copyListWithRand, input list %v", head))
	r := copyListWithRand1(head)
	if fmt.Sprint(r) == fmt.Sprint(want) {
		t.Log("copyListWithRand1 is ok")
	} else {
		t.Error("copyListWithRand1 is not ok, result:", r)
	}
	r = copyListWithRand2(head)
	if fmt.Sprint(r) == fmt.Sprint(want) {
		t.Log("copyListWithRand2 is ok")
	} else {
		t.Error("copyListWithRand2 is not ok, result:", r)
	}
}
