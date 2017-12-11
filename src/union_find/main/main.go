package main

import (
	"union_find/uf"
	"fmt"
)

//func main() {
//	var err error
//	var in, p, q int
//	var mu uf.UnionFindList
//
//	input := bufio.NewScanner(os.Stdin) //初始化一个扫表对象
//	for input.Scan() { //扫描输入内容
//		line := input.Text() //把输入内容转换为字符串
//		fmt.Println(line,"---")
//		if line == "" {
//			break
//		}
//		if !strings.Contains(line, " ") {
//			in, err = strconv.Atoi(line)
//			checkErr(err)
//			mu = uf.NewUnionFindList(in)
//			continue
//		}
//		ss := strings.Split(line, " ")
//		if ss[0]=="" || ss[1]==""{
//			fmt.Println("p or q is empty,escaped")
//			continue
//		}
//		p, err = strconv.Atoi(ss[0])
//		checkErr(err)
//		q, err = strconv.Atoi(ss[1])
//		checkErr(err)
//		mu.Connect(p, q)
//		line=""
//	}
//	fmt.Println(mu)
//
//}

func main() {
	var err error
	var in, p, q int
	var mu uf.UnionFindList

	_, err = fmt.Scanln(&in)
	checkErr(err)
	if in <= 0 {
		panic("count is invalid")
	}
	mu=uf.NewUnionFindList(in)
	for { //扫描输入内容
		_, err = fmt.Scanln(&p, &q)
		checkErr(err)
		if p == 0 && q == 0 {
			break
		}
		mu.Connect(p, q)
	}
	fmt.Println(mu)

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
