package mergesort

import "sort"

const CUTSET  = 7
type MergeSorter interface {
	sort.Interface
}

func Sort(m MergeSorter)  {

}

func _sort(src MergeSorter,dst  MergeSorter,lo,hi int)  {
	if hi<=lo+CUTSET{

	}
	return
}

func _merge(src MergeSorter,dst  MergeSorter,lo,hi int)  {

}

func insertSort(m MergeSorter,lo,hi int)  {
	for i:=lo;i<hi;i++{
		for j:=i;j>lo && m.Less(j,i);j++{
			m.Swap(i,j)
		}
	}
}