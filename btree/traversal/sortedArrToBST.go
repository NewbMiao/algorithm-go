package traversal

import (
	"github.com/NewbMiao/algorithm-go/kit/btree"
)

func SortedArrToBST(arr []int) *btree.Node {
	if len(arr) == 0 {
		return nil
	}
	return generateBST(arr, 0, len(arr)-1)
}

func generateBST(arr []int, start, end int) *btree.Node {
	if start > end {
		return nil
	}

	mid := (start + end) / 2
	head := &btree.Node{Value: arr[mid]}
	head.Left = generateBST(arr, start, mid-1)
	head.Right = generateBST(arr, mid+1, end)

	return head
}
