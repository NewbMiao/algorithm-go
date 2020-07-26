package traversal

import (
	"fmt"
	"testing"

	"github.com/NewbMiao/algorithm-go/kit/btree"
)

var cbt *btree.Node

func init() {
	cbt = &btree.Node{Value: 1}
	cbt.Left = &btree.Node{Value: 2}
	cbt.Right = &btree.Node{Value: 3}
	cbt.Left.Right = &btree.Node{Value: 4}
	cbt.Right.Left = &btree.Node{Value: 5}
	cbt.Right.Right = &btree.Node{Value: 6}
	cbt.Left.Right.Left = &btree.Node{Value: 7}
	cbt.Left.Right.Right = &btree.Node{Value: 8}
	cbt.Right.Left.Left = &btree.Node{Value: 9}
	cbt.Right.Left.Right = &btree.Node{Value: 10}
	cbt.Left.Right.Right.Right = &btree.Node{Value: 11}
	cbt.Right.Left.Left.Left = &btree.Node{Value: 12}
	cbt.Left.Right.Right.Right.Left = &btree.Node{Value: 13}
	cbt.Left.Right.Right.Right.Right = &btree.Node{Value: 14}
	cbt.Right.Left.Left.Left.Left = &btree.Node{Value: 15}
	cbt.Right.Left.Left.Left.Right = &btree.Node{Value: 16}
}

func TestEdgeMap(t *testing.T) {
	h := btree.GetHeight(cbt, 0)
	fmt.Printf("get btree height: %d,\nedge map:\n", h)
	m := make(map[int][2]*btree.Node)
	btree.GetEdgeMap(cbt, m, 0)
	for i := 0; i < h; i++ {
		fmt.Println(i, ": ", m[i][0].Value, m[i][1].Value)
	}
}

//go test  -run TestCounterClockWise -v src/btree/traversal/*
func TestCounterClockWise(t *testing.T) {
	want := []int{
		1, 2, 4, 7, 11, 13, 14, 15, 16, 12, 10, 6, 3,
	}
	t.Log(fmt.Sprintf("input btree=%v, CounterClockWise want: %v", cbt, want))
	r := CounterClockWise(cbt)
	if fmt.Sprintf("%v", r) == fmt.Sprintf("%v", want) {
		t.Log("CounterClockWise is ok")
	} else {
		t.Error("CounterClockWise is not ok, result:", fmt.Sprintf("%v", r))
	}
}

func TestCounterClockWise2(t *testing.T) {
	want := []int{
		1, 2, 4, 7, 13, 14, 15, 16, 10, 6, 3,
	}
	t.Log(fmt.Sprintf("input btree=%v, CounterClockWise2 want: %v", cbt, want))
	r := CounterClockWise2(cbt)
	if fmt.Sprintf("%v", r) == fmt.Sprintf("%v", want) {
		t.Log("CounterClockWise2 is ok")
	} else {
		t.Error("CounterClockWise2 is not ok, result:", fmt.Sprintf("%v", r))
	}
}
