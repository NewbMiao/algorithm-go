package btree

import (
	"fmt"
	"math"
	"strings"
	"unicode/utf8"
)

type Node struct {
	Left, Right *Node
	Value       int
}

func GetHeight(bt *Node, l int) int {
	if bt == nil {
		return l
	}
	return int(math.Max(float64(GetHeight(bt.Left, l+1)), float64(GetHeight(bt.Right, l+1))))
}

//获取每层的最左和最右节点
func GetEdgeMap(h *Node, edgeMap map[int][2]*Node, l int) {
	if h == nil {
		return
	}
	if edgeMap == nil {
		height := GetHeight(h, 0)
		edgeMap = make(map[int][2]*Node, height)
	}

	//最左节点
	left := edgeMap[l][0]
	if edgeMap[l][0] == nil {
		left = h
	}
	//最右节点
	right := h

	edgeMap[l] = [2]*Node{left, right}
	GetEdgeMap(h.Left, edgeMap, l+1)
	GetEdgeMap(h.Right, edgeMap, l+1)
}

func PrintTree(bt *Node) {
	fmt.Println("BinaryTree:")
	printer(bt, 0, 17, "H")
}

func printer(bt *Node, height, len int, direction string) {
	if bt == nil {
		return
	}
	printer(bt.Right, height+1, len, "v")
	val := fmt.Sprint(direction, bt.Value, direction)
	lenM := utf8.RuneCountInString(val)
	lenL := (len - lenM) / 2
	lenR := len - lenM - lenL

	str := fmt.Sprint(strings.Repeat(" ", lenL), val, strings.Repeat(" ", lenR))
	fmt.Println(strings.Repeat(" ", len*height) + str)
	printer(bt.Left, height+1, len, "^")
}
