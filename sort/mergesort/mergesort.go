package mergesort

import (
	"sort"
	"tool"
)

type MItems []string
type MergeSorter interface {
	sort.Interface
}

func Sort(m MItems) MItems {
	tool.Trace("to sort slice %v\n", m)
	n := len(m)
	cutset := 7
	if n <= cutset {
		return insertSort(m)

	}
	tool.Trace("to sort %d: %v,%v\n", n, m[:n/2], m[n/2:])
	return merge(Sort(m[:n/2]), Sort(m[n/2:]))
}

func SortBU(m MItems) MItems {
	tool.Trace("to sort slice %v\n", m)
	n := len(m)
	for step := 1; step < n; step *= 2 {
		for lo := 0; lo < n-step; lo += step * 2 {
			mid := lo + step
			hi := mid + step
			if n < hi {
				hi = n
			}
			tool.Trace("sort bottom up (%d):  %d - %d - %d\n", step, lo, mid, hi)
			tmp:=merge(m[lo:mid], m[mid:hi])
			tool.Trace("sort bottom up (%v)\n", tmp)

			for i:=lo;i<hi;i++{
				m[i]=tmp[i-lo]
			}
		}
	}
	return m
}

func merge(l, r MItems) (d MItems) {
	ln := len(l)
	rn := len(r)
	d = make(MItems, ln+rn)
	if ln == 0 || rn == 0 {
		return
	}
	tool.Trace("merge l:%v,r:%v\n", l, r)
	i := 0
	j := 0
	if l[ln-1] <= r[0] {
		tool.Trace("merge l and r once: %v,%v\n", l, r)
		d = append(l, r...)
		return
	}
	for k := 0; k < ln+rn; k++ {
		tool.Trace("to merge i:%d,j:%d\n", i, j)

		if i >= ln && j < rn {
			tool.Trace("merge right:%v\n", r[j:])
			d = append(d[:k], r[j:]...)
			break
		} else if j >= rn && i < ln {
			tool.Trace("merge left:%v\n", l[i:])
			d = append(d[:k], l[i:]...)
			break
		} else if l[i] > r[j] {
			d[k] = r[j]
			j++
		} else {
			d[k] = l[i]
			i++
		}
	}
	tool.Trace("after merge: %v\n", d)
	return
}

func insertSort(m MItems) MItems {
	for i, n := 1, len(m); i < n; i++ {
		for j := i; j > 0 && m[j] < m[j-1]; j-- {
			tool.Trace("insertSort swap %d(%v) <=> %d(%v)\n", j-1, m[j-1], j, m[j])
			m[j-1], m[j] = m[j], m[j-1]
		}
	}
	return m
}
