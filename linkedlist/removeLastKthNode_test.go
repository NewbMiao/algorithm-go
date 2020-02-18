package linkedlist

import (
	"testing"
	"fmt"
	. "github.com/NewbMiao/AlgorithmCodeNote/kit/linkedlist"
)

func TestRemoveLastKthNode(t *testing.T) {
	wants := [][]int{
		{2, 3, 4, 5, 6},
		{1, 2, 3, 5, 6},
		{1, 2, 3, 4, 5, 6},
	}

	kth := []int{6, 3, 7}
	for k := range wants {
		h1 := NewList()
		h1.Push(1)
		h1.Push(2)
		h1.Push(3)
		h1.Push(4)
		h1.Push(5)
		h1.Push(6)
		t.Log(fmt.Sprintf("RemoveLastKthNode, input list %v, remove last %dth node, want: %v", h1, kth[k], wants[k]))
		RemoveLastKthNode(h1, kth[k])
		if fmt.Sprint(h1) == fmt.Sprint(wants[k]) {
			t.Log("RemoveLastKthNode is ok")
		} else {
			t.Error("RemoveLastKthNode is not ok, result:", fmt.Sprintf("%v", h1))
		}
	}

}

func TestRemoveLastKthDNode(t *testing.T) {
	wants := [][]int{
		{2, 3, 4, 5, 6},
		{1, 2, 3, 5, 6},
		{1, 2, 3, 4, 5, 6},
	}

	kth := []int{6, 3, 7}
	for k := range wants {
		h1 := NewDList()
		h1.PushBack(1)
		h1.PushBack(2)
		h1.PushBack(3)
		h1.PushBack(4)
		h1.PushBack(5)
		h1.PushBack(6)
		t.Log(fmt.Sprintf("RemoveLastKthDNode, input list %v, remove last %dth node, want: %v", h1, kth[k], wants[k]))
		RemoveLastKthDNode(h1, kth[k])
		if fmt.Sprint(h1) == fmt.Sprint(wants[k]) {
			t.Log("RemoveLastKthDNode is ok")
		} else {
			t.Error("RemoveLastKthDNode is not ok, result:", fmt.Sprintf("%v", h1))
		}
	}

}
