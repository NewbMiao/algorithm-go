package uf

import (
	"errors"
	"fmt"
	"io"

	"github.com/NewbMiao/algorithm-go/tool"
	"github.com/NewbMiao/algorithm-go/union_find/uf"
)

func Scanln() {
	var err error
	var in, p, q, re int
	var mu uf.UnionFindList
	isEoF := false
	_, err = fmt.Scanln(&in)
	tool.ErrPanic(err)
	mu = uf.NewUnionFindList(in)

	for !isEoF { // 扫描输入内容
		re, err = fmt.Scanln(&p, &q)
		if errors.Is(err, io.EOF) {
			err = nil
			isEoF = true
		}
		if re != 2 {
			break
		}
		tool.ErrPanic(err)
		mu.Connect(p, q)
	}
	fmt.Println(mu, mu.Count())
}
