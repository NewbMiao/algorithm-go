package subtree

import (
	"fmt"
	"testing"

	"github.com/NewbMiao/algorithm-go/kit/btree"
)

var ebt, ebt1, ebt2 *btree.Node

func init() {
	ebt = &btree.Node{Value: 6}
	ebt.Left = &btree.Node{Value: 3}
	ebt.Right = &btree.Node{Value: 7}
	ebt.Left.Left = &btree.Node{Value: 2}
	ebt.Left.Right = &btree.Node{Value: 4}
	ebt.Right.Left = &btree.Node{Value: 5}
	ebt.Right.Right = &btree.Node{Value: 8}
	ebt.Left.Left.Left = &btree.Node{Value: 1}

	ebt1 = &btree.Node{Value: 8}
	ebt1.Left = &btree.Node{Value: 3}
	ebt1.Right = &btree.Node{Value: 7}
	ebt1.Left.Left = &btree.Node{Value: 2}
	ebt1.Left.Right = &btree.Node{Value: 4}
	ebt1.Right.Left = &btree.Node{Value: 6}
	ebt1.Right.Right = &btree.Node{Value: 5}
	ebt1.Left.Left.Left = &btree.Node{Value: 1}

	ebt2 = &btree.Node{Value: 3}
	ebt2.Left = &btree.Node{Value: 5}
	ebt2.Right = &btree.Node{Value: 7}
	ebt2.Left.Left = &btree.Node{Value: 2}
	ebt2.Left.Right = &btree.Node{Value: 4}
	ebt2.Right.Left = &btree.Node{Value: 6}
	ebt2.Right.Right = &btree.Node{Value: 8}
	ebt2.Left.Left.Left = &btree.Node{Value: 1}
}

func TestFindTwoErrNode(t *testing.T) {
	wants := map[*btree.Node][]int{
		ebt:  {6, 5},
		ebt1: {8, 5},
		ebt2: {5, 3},
	}
	for bt, want := range wants {
		t.Log(fmt.Sprintf("FindTwoErrNode  want: %v", want))
		btree.PrintTree(bt)
		r := FindTwoErrNode(bt)
		p := FindTwoErrNodeParent(bt, r)
		tmp := []int{}
		tmpP := []int{}
		for _, v := range r {
			tmp = append(tmp, v.Value)
		}
		for _, v := range p {
			if v == nil {
				tmpP = append(tmpP, 0)
				continue
			}
			tmpP = append(tmpP, v.Value)
		}
		fmt.Println("errNode:", tmp, "parentNode:", tmpP)
		btree.PrintTree(RecoverTree(bt))
		if fmt.Sprintf("%v", want) == fmt.Sprintf("%v", tmp) {
			t.Log("FindTwoErrNode is ok")
		} else {
			t.Error("FindTwoErrNode is not ok, result:", fmt.Sprintf("%v", tmp))
		}
	}
}
