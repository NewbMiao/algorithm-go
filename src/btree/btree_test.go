package btree

import (
	"testing"
	"math"
	"fmt"
)

var gbt *Node

func init() {
	gbt = &Node{Value: 1}
	gbt.Left = &Node{Value: 2}
	gbt.Right = &Node{Value: 3}
	gbt.Left.Left = &Node{Value: 4}
	gbt.Right.Left = &Node{Value: 5}
	gbt.Right.Right = &Node{Value: 6}
	gbt.Left.Left.Right = &Node{Value: 7}
}
func TestPrinter(t *testing.T) {
	bt := &Node{Value: 1}
	bt.Left = &Node{Value: -222222222}
	bt.Right = &Node{Value: 3}
	bt.Left.Left = &Node{Value: math.MinInt32}
	bt.Right.Left = &Node{Value: 55555555}
	bt.Right.Right = &Node{Value: 66}
	bt.Left.Left.Right = &Node{Value: 777}
	PrintTree(bt)

	bt = &Node{Value: 1}
	bt.Left = &Node{Value: 1}
	bt.Right = &Node{Value: 1}
	bt.Left.Left = &Node{Value: 1}
	bt.Right.Left = &Node{Value: 1}
	bt.Right.Right = &Node{Value: 1}
	bt.Left.Left.Right = &Node{Value: 1}
	PrintTree(bt)

	PrintTree(gbt)

}

func TestEdgeMap(t *testing.T) {
	h := GetHeight(gbt, 0)
	fmt.Printf("get gbtree height: %d \nedge map:\n", h)
	m := make(map[int][2]*Node)
	GetEdgeMap(gbt, m, 0)
	for i := 0; i < h; i++ {
		fmt.Println(i, ": ", m[i][0].Value, m[i][1].Value)
	}
}
