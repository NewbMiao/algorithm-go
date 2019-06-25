package traversal

import (
	"testing"
	"fmt"
	. "btree"
)

var mbt *Node

func init() {
	mbt = &Node{Value: 4}
	mbt.Left = &Node{Value: 2}
	mbt.Right = &Node{Value: 6}
	mbt.Left.Left = &Node{Value: 1}
	mbt.Left.Right = &Node{Value: 3}
	mbt.Right.Left = &Node{Value: 5}
	mbt.Right.Right = &Node{Value: 7}
	PrintTree(mbt)

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
	want := []int{1,3,2,5,7,6,4}
	t.Log(fmt.Sprintf("MorrisLast want: %v", want))
	r := MorrisLast(mbt)
	if fmt.Sprintf("%v", r) == fmt.Sprintf("%v", want) {
		t.Log("MorrisLast is ok")
	} else {
		t.Error("MorrisLast is not ok, result:", r)
	}
}
