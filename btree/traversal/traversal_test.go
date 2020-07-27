package traversal

import (
	"fmt"
	"testing"

	"github.com/NewbMiao/algorithm-go/kit/btree"
)

var tbt *btree.Node

func init() {
	tbt = &btree.Node{Value: 5}
	tbt.Left = &btree.Node{Value: 3}
	tbt.Right = &btree.Node{Value: 8}
	tbt.Left.Left = &btree.Node{Value: 2}
	tbt.Left.Right = &btree.Node{Value: 4}
	tbt.Left.Left.Left = &btree.Node{Value: 1}
	tbt.Right.Left = &btree.Node{Value: 7}
	tbt.Right.Left.Left = &btree.Node{Value: 6}
	tbt.Right.Right = &btree.Node{Value: 10}
	tbt.Right.Right.Left = &btree.Node{Value: 9}
	tbt.Right.Right.Right = &btree.Node{Value: 11}
	btree.PrintTree(tbt)
}

func TestPreOrderRecur(t *testing.T) {
	want := []int{
		5, 3, 2, 1, 4, 8, 7, 6, 10, 9, 11,
	}
	t.Log(fmt.Sprintf("pre order want: %v", want))
	r := PreOrderRecur(tbt)
	if fmt.Sprintf("%v", r) == fmt.Sprintf("%v", want) {
		t.Log("pre order is ok")
	} else {
		t.Error("pre order is not ok, result:", fmt.Sprintf("%v", r))
	}
}

func TestInOrderRecur(t *testing.T) {
	want := []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	}
	t.Log(fmt.Sprintf("in order want: %v", want))
	r := InOrderRecur(tbt)
	if fmt.Sprintf("%v", r) == fmt.Sprintf("%v", want) {
		t.Log("in order is ok")
	} else {
		t.Error("in order is not ok, result:", fmt.Sprintf("%v", r))
	}
}

func TestLastOrderRecur(t *testing.T) {
	want := []int{
		1, 2, 4, 3, 6, 7, 9, 11, 10, 8, 5,
	}
	t.Log(fmt.Sprintf("last order want: %v", want))
	r := LastOrderRecur(tbt)
	if fmt.Sprintf("%v", r) == fmt.Sprintf("%v", want) {
		t.Log("last order is ok")
	} else {
		t.Error("last order is not ok, result:", fmt.Sprintf("%v", r))
	}
}

func TestPreOrder(t *testing.T) {
	want := []int{
		5, 3, 2, 1, 4, 8, 7, 6, 10, 9, 11,
	}
	t.Log(fmt.Sprintf("pre order want: %v", want))
	r := PreOrder(tbt)
	if fmt.Sprintf("%v", r) == fmt.Sprintf("%v", want) {
		t.Log("pre order is ok")
	} else {
		t.Error("pre order is not ok, result:", fmt.Sprintf("%v", r))
	}
}

func TestInOrder(t *testing.T) {
	want := []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	}
	t.Log(fmt.Sprintf("in order want: %v", want))
	r := InOrder(tbt)
	if fmt.Sprintf("%v", r) == fmt.Sprintf("%v", want) {
		t.Log("in order is ok")
	} else {
		t.Error("in order is not ok, result:", fmt.Sprintf("%v", r))
	}
}

func TestLastOrder1(t *testing.T) {
	want := []int{
		1, 2, 4, 3, 6, 7, 9, 11, 10, 8, 5,
	}
	t.Log(fmt.Sprintf("last order1 want: %v", want))
	r := LastOrder1(tbt)
	if fmt.Sprintf("%v", r) == fmt.Sprintf("%v", want) {
		t.Log("last order1 is ok")
	} else {
		t.Error("last order1 is not ok, result:", fmt.Sprintf("%v", r))
	}
}

func TestLastOrder2(t *testing.T) {
	want := []int{
		1, 2, 4, 3, 6, 7, 9, 11, 10, 8, 5,
	}
	t.Log(fmt.Sprintf("last order2 want: %v", want))
	r := LastOrder2(tbt)
	if fmt.Sprintf("%v", r) == fmt.Sprintf("%v", want) {
		t.Log("last order2 is ok")
	} else {
		t.Error("last order2 is not ok, result:", fmt.Sprintf("%v", r))
	}
}
