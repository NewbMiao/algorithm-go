package main

import (
	"union_find/uf"
	"fmt"
	"strings"
	"bufio"
	"os"
	"strconv"
	"tool"
)

func main() {
	var err error
	var in, p, q, re int
	var mu uf.UnionFindList

	input := bufio.NewScanner(os.Stdin) //初始化一个扫表对象
	for input.Scan() { //扫描输入内容
		if err := input.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}
		line := input.Text() //把输入内容转换为字符串
		if line==""{
			break
		}
		if !strings.Contains(line, " ") {
			in, err = strconv.Atoi(line)
			tool.ErrPanic(err)
			mu = uf.NewUnionFindList(in)
			continue
		}
		re, err = fmt.Sscanf(line, "%d %d", &p, &q)
		tool.ErrPanic(err)

		if re != 2 {
			fmt.Println("p or q is empty,escaped")
			continue
		}

		mu.Connect(p, q)
	}
	//fmt.Println(mu)

	fmt.Println(mu.Count())

}
