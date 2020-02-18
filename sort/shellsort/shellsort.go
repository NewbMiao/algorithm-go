package shellsort

import (
	"sort"
	"fmt"
)

type ShellSorter interface {
	sort.Interface
}

func Sort(arr ShellSorter) {

	h := 1
	for h < 3*h { // find a seq by 3x+1
		h = 3*h + 1
	}
	n := arr.Len()
	for h >= 1 {
		for i := h; i < n; i += 1 {
			for j := i; j >= h && arr.Less(j, j-h); j -= h {
				arr.Swap(j, j-h)
				fmt.Printf("swap %d,%d %v\n", j, j-h, arr)

			}
		}
		h /= 3
	}
}
