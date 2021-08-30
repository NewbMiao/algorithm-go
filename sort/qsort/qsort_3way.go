package qsort

/*
* Three Way Partitioning

* By Bently and McIlroy

*
* left part center part right part

* +-------------------------------------------------------+

* lo| = pivot | < pivot | ? | > pivot | = pivot |

* +-------------------------------------------------------+

*             ^         ^   ^         ^

*             |         |   |         |

*            less->       i-> <-j   <-great

*.
 */
func _qsort3way(arr Sitem, lo, hi int) {
	N := hi - lo + 1
	if N <= CUTOFF {
		insertSort(arr, lo, hi)
		return
	}
	// Bentley-McIlroy 3-way partitioning
	i, j := lo, hi+1
	less, great := lo, hi+1
	for {
		for {
			i++
			if !arr.Less(i, lo) || i == hi {
				break
			}
		}
		for {
			j--
			if !arr.Less(lo, j) || j == lo {
				break
			}
		}
		// pointers cross
		if i == j && arr[i] == arr[lo] {
			less++
			arr.Swap(less, i)
		}
		if i >= j {
			break
		}
		arr.Swap(i, j)

		// 这里暂时将与pivot相等的元素换到数组的两端
		if arr[i] == arr[lo] {
			less++
			arr.Swap(less, i)
		}
		if arr[j] == arr[lo] {
			great--
			arr.Swap(great, j)
		}
	}

	// 将相等的元素交换到中间
	i = j + 1

	for k := lo; k <= less; k++ {
		arr.Swap(k, j)
		j--
	}
	for k := hi; k >= great; k-- {
		arr.Swap(k, i)
		i++
	}

	_qsort3way(arr, lo, j)
	_qsort3way(arr, i, hi)
}

func Sort3way(arr Sitem) {
	lo, hi := 0, arr.Len()-1
	_qsort3way(arr, lo, hi)
}
