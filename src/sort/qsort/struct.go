package qsort

import "tool"

const CUTOFF = 8

type Sitem []string

func (s Sitem) Len() int {
	return len(s)
}

func (s Sitem) Less(i, j int) bool {
	return s[i] < s[j]
}
func (s Sitem) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func insertSort(m Sitem, lo, hi int) Sitem {
	for i := lo; i <= hi; i++ {
		for j := i; j > 0 && m[j] < m[j-1]; j-- {
			tool.Trace("swap %d(%v) <=> %d(%v)\n", j-1, m[j-1], j, m[j])
			m[j-1], m[j] = m[j], m[j-1]
		}
	}
	return m
}
