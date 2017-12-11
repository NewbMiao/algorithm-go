//https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/UF.java.html
package uf

import (
	"fmt"
	"errors"
)

type UnionFindList []int

var count int
var rank []byte

func NewUnionFindList(n int) UnionFindList {
	if n <= 0 {
		panic(errors.New("count is not postive num"))
	}
	count = n
	mu := make(UnionFindList, count)
	rank = make([]byte, count)
	for k := range mu {
		mu[k] = k
	}
	return mu
}

func (u *UnionFindList) find(i int) int {
	if i >= len(*u) {
		panic(fmt.Errorf("index out of range"))
	}
	ti := (*u)[i]
	for ti != i {
		ti = (*u)[ti]
	}
	return ti
}
func (u *UnionFindList) IsConnected(p, q int) bool {
	if u.find(p) == u.find(q) {
		return true
	}
	return false
}

func (u *UnionFindList) Connect(p, q int) {
	tp, tq := u.find(p), u.find(q)
	if tp == tq {
		return
	}
	if rank[tp] < rank[tq] {
		(*u)[tp] = (*u)[tq]
	} else if rank[tp] > rank[tq] {
		(*u)[tq] = (*u)[tp]
	} else {
		(*u)[tq] = (*u)[tp]
		rank[tp]++
	}

}

func (u *UnionFindList) Count(p, q int) int {
	return count
}

func (u *UnionFindList) String() {
	fmt.Println(fmt.Sprintf("%v", u))
}
