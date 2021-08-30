package uf

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/NewbMiao/algorithm-go/tool"
	"github.com/NewbMiao/algorithm-go/union_find/uf"
)

func BufReader() {
	var err error
	var in, p, q int
	var mu uf.UnionFindList
	isEoF := false
	var line string
	input := bufio.NewReader(os.Stdin) // 初始化一个扫表对象
	for !isEoF {                       // 扫描输入内容
		line, err = input.ReadString('\n')
		if errors.Is(err, io.EOF) {
			isEoF = true
			err = nil
		}
		if line == "\n" || line == "" {
			break
		}
		if err != nil && !errors.Is(err, io.EOF) {
			break
		}
		tool.ErrPanic(err)
		var re int
		if !strings.Contains(line, " ") {
			_, err = fmt.Sscanf(line, "%d", &in)
			tool.ErrPanic(err)
			mu = uf.NewUnionFindList(in)
			continue
		}
		re, err = fmt.Sscanf(line, "%d %d\n", &p, &q)
		tool.ErrPanic(err)

		if re != 2 {
			fmt.Println("p or q is empty,escaped")
			continue
		}
		mu.Connect(p, q)
	}

	// fmt.Println(mu)
	fmt.Println(mu.Count())
}
