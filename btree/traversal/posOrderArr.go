package traversal

import (
	"github.com/NewbMiao/algorithm-go/kit/btree"
)

//BST post order arr check
func IsPosOrderArr(arr []int) bool {
	if len(arr) == 0 {
		return false
	}
	return isPosOrder(arr, 0, len(arr)-1)
}

func isPosOrder(arr []int, start, end int) bool {
	if start == end { //self no need cmp
		return true
	}
	less := -1
	more := end
	for i := start; i < end; i++ {
		if arr[end] > arr[i] {
			less = i
		} else if more == end {
			more = i
		}
	}
	if less == -1 || more == end {
		return isPosOrder(arr, start, end-1)
	}
	if less != more-1 { //no asc order
		return false
	}

	return isPosOrder(arr, start, less) && isPosOrder(arr, more, end-1)
}

func PosOrderArrToBST(arr []int) (bt *btree.Node) {
	if len(arr) == 0 {
		return
	}
	return posToBST(arr, 0, len(arr)-1)
}

func posToBST(arr []int, start, end int) (bt *btree.Node) {
	if start > end { //should finish current loop
		return
	}
	less := -1
	more := end
	head := &btree.Node{Value: arr[end]}
	for i := start; i < end; i++ {
		if arr[end] > arr[i] {
			less = i
		} else if more == end {
			more = i
		}
	}
	head.Left = posToBST(arr, start, less)
	head.Right = posToBST(arr, more, end-1)

	return head
}
