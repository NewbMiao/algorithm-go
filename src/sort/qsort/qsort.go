package qsort

import (
	"sort"
	"tool"
)

func qsort(arr QuickSorter, lo, hi int) {
	if hi <= lo {
		return
	}
	index := partition(arr, lo, hi)
	qsort(arr, lo, index-1)
	qsort(arr, index+1, hi)
}

func Sort(arr QuickSorter) {
	lo, hi := 0, arr.Len()-1
	qsort(arr, lo, hi)
}

type QuickSorter interface {
	sort.Interface
}

func partition(arr QuickSorter, lo, hi int) int {
	i, j := lo+1, hi
	for {
		for arr.Less(i, lo) {
			if i == hi {
				break
			}
			i++
		}
		for arr.Less(lo, j) {
			if j == lo {
				break
			}
			j--
		}

		if i >= j {
			break
		}
		tool.Trace("partition swapping %d,%d\n", lo, j)
		arr.Swap(i, j)
	}
	tool.Trace("partition last swap %d,%d\n", lo, j)
	arr.Swap(lo, j)
	return j

}
