package qsort

import (
	"sort"
	"tool"
)

func qSort(arr QuickSorter, lo, hi int) {
	if hi <= lo {
		return
	}
	index := partition(arr, lo, hi)
	qSort(arr, lo, index-1)
	qSort(arr, index+1, hi)
}

func Sort(arr QuickSorter) {
	lo, hi := 0, arr.Len()-1
	qSort(arr, lo, hi)
}

type QuickSorter interface {
	sort.Interface
}

func partition(arr QuickSorter, lo, hi int) int {
	i, j := lo+1, hi
	for {
		for arr.Less(i, lo) && i < hi {
			i++
		}
		for arr.Less(lo, j) { //无需判断 j > lo，因为 j=lo不满足less
			j--
		}

		if i >= j {
			break
		}
		tool.Trace("partition swapping %d,%d\n", i, j)
		arr.Swap(i, j)
	}
	tool.Trace("partition last swap %d,%d\n", lo, j)
	arr.Swap(lo, j)
	return j

}
