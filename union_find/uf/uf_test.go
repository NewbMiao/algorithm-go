package uf

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/NewbMiao/algorithm-go/tool"
)

func TestUF(t *testing.T) {
	mu := readBuf()
	if mu.Count() == 2 {
		t.Log("Component count is OK: 2")
	} else {
		t.Error("Component count is Not OK")
	}
	want := "[6 6 6 4 4 6 6 6 4 4]"
	if mu.String() == want {
		t.Log("Union Find List is OK:", want)
	} else {
		t.Error("Union Find List is Not OK: ", mu.String(), ",Want:", want)
	}
}
func readBuf() (mu UnionFindList) {
	var err error
	var in, p, q, re int
	//tinyUF.txt
	str := `10
4 3
3 8
6 5
9 4
2 1
8 9
5 0
7 2
6 1
1 0
6 7
`
	f := strings.NewReader(str)
	input := bufio.NewScanner(f)

	for input.Scan() { //扫描输入内容
		if err := input.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}
		line := input.Text() //把输入内容转换为字符串
		if !strings.Contains(line, " ") {
			in, err = strconv.Atoi(line)
			tool.ErrPanic(err)
			fmt.Println(in)
			mu = NewUnionFindList(in)
			continue
		}
		re, err = fmt.Sscanf(line, "%d %d", &p, &q)
		fmt.Println(p, q)
		tool.ErrPanic(err)

		if re != 2 {
			fmt.Println("p or q is empty,escaped")
			continue
		}
		mu.Connect(p, q)
	}
	return mu
}
