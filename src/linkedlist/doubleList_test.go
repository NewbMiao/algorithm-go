package linkedlist

import (
	"testing"
	"fmt"
)

func TestDList_PushBack(t *testing.T) {
	dl := NewDList()
	dl.PushBack(1)
	dl.PushBack(2)
	dl.PushBack(3)

	want := []int{1, 2, 3}
	t.Log(fmt.Sprintf("DList_PushBack, input doublelist %v, want: %v", dl, want))
	if fmt.Sprint(dl) == fmt.Sprint(want) {
		t.Log("DList_PushBack is ok")
	} else {
		t.Error("DList_PushBack is not ok, result:", fmt.Sprintf("%v", dl))
	}
}

func TestDList_PushFront(t *testing.T) {
	dl := NewDList()
	dl.PushFront(1)
	dl.PushFront(2)
	dl.PushFront(3)

	want := []int{3, 2, 1}
	t.Log(fmt.Sprintf("DList_PushFront, input doublelist %v, want: %v", dl, want))
	if fmt.Sprint(dl) == fmt.Sprint(want) {
		t.Log("DList_PushFront is ok")
	} else {
		t.Error("DList_PushFront is not ok, result:", fmt.Sprintf("%v", dl))
	}
}

func TestDList_PopBack(t *testing.T) {
	dl := NewDList()
	dl.PushBack(1)
	dl.PushBack(2)
	dl.PushBack(3)

	dl.PopBack()
	want := []int{1, 2}
	t.Log(fmt.Sprintf("DList_PopBack, input doublelist %v, want: %v", dl, want))
	if fmt.Sprint(dl) == fmt.Sprint(want) {
		t.Log("DList_PopBack is ok")
	} else {
		t.Error("DList_PopBack is not ok, result:", fmt.Sprintf("%v", dl))
	}
}

func TestDList_PopFront(t *testing.T) {
	dl := NewDList()
	dl.PushFront(1)
	dl.PushFront(2)
	dl.PushFront(3)

	dl.PopFront()
	want := []int{2, 1}
	t.Log(fmt.Sprintf("DList_PopFront, input doublelist %v, want: %v", dl, want))
	if fmt.Sprint(dl) == fmt.Sprint(want) {
		t.Log("DList_PopFront is ok")
	} else {
		t.Error("DList_PopFront is not ok, result:", fmt.Sprintf("%v", dl))
	}
}

func TestDList_IsEmpty(t *testing.T) {
	dl := NewDList()
	fmt.Println(dl, "dl init, isEmpty", dl.IsEmpty(), dl.Size)
	dl.PushBack(1)
	dl.PopBack()
	fmt.Println(dl, "dl push and pop back, isEmpty", dl.IsEmpty(), dl.Size)

	dl.PushFront(1)
	dl.PopFront()
	fmt.Println(dl, "dl push and pop front, isEmpty", dl.IsEmpty(), dl.Size)
	dl.PushFront(1)
	dl.PushFront(2)
	fmt.Println(dl, "front:", dl.Front(), "back:", dl.Back())

	r1 := dl.PopFront()
	r2 := dl.PopBack()
	r3 := dl.PopFront()

	fmt.Println(dl, "pop all:",r1, r2, r3, "isEmpty", dl.IsEmpty(), dl.Size)
	fmt.Println(dl, "dl is empty, isLast", dl.IsLast(dl.Front()), dl.Size)

	dl.PushFront(1)
	fmt.Println(dl, "dl push front, isLast", dl.IsLast(dl.Front()), dl.Size)
}
