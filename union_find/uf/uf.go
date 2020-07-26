//https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/UF.java.html
package uf

import (
	"errors"
	"fmt"

	"github.com/NewbMiao/algorithm-go/tool"
)

type UnionFindList []int

var count int
var rank []byte

func NewUnionFindList(n int) UnionFindList {
	if n <= 0 {
		panic(errors.New("count is invalid"))
	}
	tool.Trace("init union find list, size:%d\n", n)

	count = n
	mu := make(UnionFindList, count)
	rank = make([]byte, count)
	for k := range mu {
		mu[k] = k
	}
	return mu
}

func (parent *UnionFindList) find(i int) int {
	if i >= len(*parent) || i < 0 {
		panic(fmt.Errorf("index not between 0 and %d", len(*parent)))
	}
	tool.Trace("\tstart find the root of %d: \n", i)
	for i != (*parent)[i] {
		tool.Trace("\tfinding: index %d,parent %d\n", i, (*parent)[i])
		(*parent)[i] = (*parent)[(*parent)[i]] //path compression by halving
		tool.Trace("\tpath compression by halving: chg index %d's parent to index %d\n", i, (*parent)[i])

		i = (*parent)[(*parent)[i]]
	}
	return i
}
func (parent *UnionFindList) IsConnected(p, q int) bool {
	tool.Trace("check %d & %d is connected\n", p, q)
	if parent.find(p) == parent.find(q) {
		tool.Trace("already connected: %d & %d\n", p, q)
		return true
	}
	return false
}

func (parent *UnionFindList) Connect(p, q int) {
	tool.Trace("try to connect %d & %d\n", p, q)
	rootP, rootQ := parent.find(p), parent.find(q)
	if rootP == rootQ {
		tool.Trace("no need connect: %d & %d\n", p, q)
		return
	}

	//move little rank to larger rank's root
	if rank[rootP] < rank[rootQ] {
		(*parent)[rootP] = rootQ
	} else if rank[rootP] > rank[rootQ] {
		(*parent)[rootQ] = rootP
	} else {
		(*parent)[rootQ] = rootP
		rank[rootP]++
	}
	//after connected,node is decrease one
	count--
	tool.Trace("after connect: \n\tcount:\t%d\n\tparent:\t%v\n\trank:\t%v\n", count, *parent, rank)
}

func (parent *UnionFindList) Count() int {
	return count
}

func (parent *UnionFindList) String() string {
	return fmt.Sprintf("%v", *parent)
}
