package qsort

import (
	"github.com/NewbMiao/AlgorithmCodeNote/tool"
)

func _qsort(arr Sitem, lo, hi int) {
	N := hi - lo + 1
	if N <= CUTOFF {
		insertSort(arr, lo, hi)
		return
	} else if N <= 40 { //use median of three
		m := median3(arr, lo, lo+N>>1, hi)
		arr.Swap(lo, m)
	} else { //use median of ninther
		eps := N / 8;
		mid := lo + N>>1;
		m1 := median3(arr, lo, lo+eps, lo+eps+eps);
		m2 := median3(arr, mid-eps, mid, mid+eps);
		m3 := median3(arr, hi-eps-eps, hi-eps, hi);
		ninther := median3(arr, m1, m2, m3);
		arr.Swap(lo, ninther)
	}
	index := _partition(arr, lo, hi)
	_qsort(arr, lo, index-1)
	_qsort(arr, index+1, hi)
}

func SortX(arr Sitem) {
	lo, hi := 0, arr.Len()-1
	_qsort(arr, lo, hi)
}

func _partition(arr Sitem, lo, hi int) int {
	i, j := lo+1, hi
	for {
		for arr.Less(i, lo) {
			if i == hi {
				break
			}
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

func median3(arr Sitem, x, y, z int) int {
	if arr[x] < arr[y] { //x < y
		if arr[y] < arr[z] { //y < z
			return y
		} else if arr[x] < arr[z] { // x< z < y
			return z
		} else { // y > x >z
			return x
		}
	} else { //x > y
		if arr[y] > arr[z] { //y > z
			return y
		} else if arr[x] > arr[z] { // x>z>y
			return z
		} else { // y<x<z
			return x
		}
	}
}
