package linkedlist

import (
	"testing"
	"fmt"
	. "kit/linkedlist"
)

func TestGetIntersectNode1(t *testing.T) {
	l1 := NewList()
	l1.Push(1)
	l1.Push(2)
	l1.Push(3)
	l1.Push(4)
	l1.Push(5)
	l1.Push(6)
	l1.Push(7)

	l2 := NewList()
	l2.Push(0)
	l2.Push(9)
	l2.Push(8)

	//l2 8-> l1 6
	l2.Head.Next.Next.Next = l1.Head.Next.Next.Next.Next.Next
	t.Log(fmt.Sprintf("getIntersectNode1, input list %v, %v", l1, l2))

	want := 6
	r := getIntersectNode(l1.Head, l2.Head)

	if r != nil && r.Value == want {
		t.Log("getIntersectNode1 is ok")
	} else {
		t.Error("getIntersectNode1 is not ok, result:", r)
	}
}

func TestGetIntersectNode2(t *testing.T) {
	l1 := NewList()
	l1.Push(1)
	l1.Push(2)
	l1.Push(3)
	l1.Push(4)
	l1.Push(5)
	l1.Push(6)
	l1.Push(7)
	l1.Head.Next.Next.Next.Next.Next.Next.Next = l1.Head.Next.Next.Next

	l2 := NewList()
	l2.Push(0)
	l2.Push(9)
	l2.Push(8)

	//l2 8-> l1 2
	l2.Head.Next.Next.Next = l1.Head.Next
	t.Log(fmt.Sprint("getIntersectNode2"))

	want := 2
	r := getIntersectNode(l1.Head, l2.Head)

	if r != nil && r.Value == want {
		t.Log("getIntersectNode2 is ok")
	} else {
		t.Error("getIntersectNode2 is not ok, result:", r)
	}
}

func TestGetIntersectNode3(t *testing.T) {
	l1 := NewList()
	l1.Push(1)
	l1.Push(2)
	l1.Push(3)
	l1.Push(4)
	l1.Push(5)
	l1.Push(6)
	l1.Push(7)
	// 7 -> 4
	l1.Head.Next.Next.Next.Next.Next.Next.Next = l1.Head.Next.Next.Next

	l2 := NewList()
	l2.Push(0)
	l2.Push(9)
	l2.Push(8)

	//l2 8-> l1 6
	l2.Head.Next.Next.Next = l1.Head.Next.Next.Next.Next.Next
	t.Log(fmt.Sprint("getIntersectNode3"))

	want := 4
	r := getIntersectNode(l1.Head, l2.Head)

	if r != nil && r.Value == want {
		t.Log("getIntersectNode3 is ok")
	} else {
		t.Error("getIntersectNode3 is not ok, result:", r)
	}
}
