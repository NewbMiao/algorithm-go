package uf

import (
	"github.com/NewbMiao/AlgorithmCodeNote/union_find/uf"
	"fmt"
	"io"
	"github.com/NewbMiao/AlgorithmCodeNote/tool"
)

func Scanln() {
	var err error
	var in, p, q, re int
	var mu uf.UnionFindList
	var isEoF = false
	re, err = fmt.Scanln(&in)
	tool.ErrPanic(err)
	mu = uf.NewUnionFindList(in)

	for !isEoF { //扫描输入内容
		re, err = fmt.Scanln(&p, &q)
		if err == io.EOF {
			err = nil
			isEoF = true
		}
		if re != 2 {
			break
		}
		tool.ErrPanic(err)
		mu.Connect(p, q)

	}
	fmt.Println(mu,mu.Count())

}
