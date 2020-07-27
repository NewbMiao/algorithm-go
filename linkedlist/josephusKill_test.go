package linkedlist

import (
	"fmt"
	"testing"

	"github.com/NewbMiao/algorithm-go/kit/linkedlist"
)

func TestJosephusKill1(t *testing.T) {
	l := linkedlist.NewList()
	l.Push(1)
	l.Push(2)
	l.Push(3)
	l.Push(4)
	l.Push(5)
	//成环
	l.Last().Next = l.Head

	m := 3
	want := 4
	t.Log(fmt.Sprintf("JosephusKill1, input list %v, m %d, want: %v", l, m, want))
	r := JosephusKill1(l, m)
	if r == want {
		t.Log("JosephusKill1 is ok")
	} else {
		t.Error("JosephusKill1 is not ok, result:", r)
	}
}

func TestJosephusKill2(t *testing.T) {
	l := linkedlist.NewList()
	l.Push(1)
	l.Push(2)
	l.Push(3)
	l.Push(4)
	l.Push(5)
	//成环
	l.Last().Next = l.Head

	m := 3
	want := 4
	t.Log(fmt.Sprintf("JosephusKill2, input list %v, m %d, want: %v", l, m, want))
	r := JosephusKill2(l, m)
	if r == want {
		t.Log("JosephusKill2 is ok")
	} else {
		t.Error("JosephusKill2 is not ok, result:", r)
	}
}
