package traversal

import (
	"fmt"
	"testing"

	"github.com/NewbMiao/algorithm-go/kit/btree"
)

var mbt *btree.Node

func init() {
	mbt = &btree.Node{Value: 4}
	mbt.Left = &btree.Node{Value: 2}
	mbt.Right = &btree.Node{Value: 6}
	mbt.Left.Left = &btree.Node{Value: 1}
	mbt.Left.Right = &btree.Node{Value: 3}
	mbt.Right.Left = &btree.Node{Value: 5}
	mbt.Right.Right = &btree.Node{Value: 7}
	btree.PrintTree(mbt)
}

func TestMorrisIn(t *testing.T) {
	want := []int{1, 2, 3, 4, 5, 6, 7}
	t.Log(fmt.Sprintf("MorrisIn want: %v", want))
	r := MorrisIn(mbt)
	if fmt.Sprintf("%v", r) == fmt.Sprintf("%v", want) {
		t.Log("MorrisIn is ok")
	} else {
		t.Error("MorrisIn is not ok, result:", r)
	}
}

func TestMorrisPre(t *testing.T) {
	want := []int{4, 2, 1, 3, 6, 5, 7}
	t.Log(fmt.Sprintf("MorrisPre want: %v", want))
	r := MorrisPre(mbt)
	if fmt.Sprintf("%v", r) == fmt.Sprintf("%v", want) {
		t.Log("MorrisPre is ok")
	} else {
		t.Error("MorrisPre is not ok, result:", r)
	}
}

func TestMorrisLast(t *testing.T) {
	want := []int{1, 3, 2, 5, 7, 6, 4}
	t.Log(fmt.Sprintf("MorrisLast want: %v", want))
	r := MorrisLast(mbt)
	if fmt.Sprintf("%v", r) == fmt.Sprintf("%v", want) {
		t.Log("MorrisLast is ok")
	} else {
		t.Error("MorrisLast is not ok, result:", r)
	}
}
