package traversal

import (
	"testing"
	. "kit/btree"
)

func TestSortedArrToBST(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	PrintTree(SortedArrToBST(data))
}
