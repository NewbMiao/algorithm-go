package subtree

import (
	"testing"
	"fmt"
	. "kit/btree"
)

var ebt, ebt1, ebt2 *Node

func init() {
	ebt = &Node{Value: 6}
	ebt.Left = &Node{Value: 3}
	ebt.Right = &Node{Value: 7}
	ebt.Left.Left = &Node{Value: 2}
	ebt.Left.Right = &Node{Value: 4}
	ebt.Right.Left = &Node{Value: 5}
	ebt.Right.Right = &Node{Value: 8}
	ebt.Left.Left.Left = &Node{Value: 1}

	ebt1 = &Node{Value: 8}
	ebt1.Left = &Node{Value: 3}
	ebt1.Right = &Node{Value: 7}
	ebt1.Left.Left = &Node{Value: 2}
	ebt1.Left.Right = &Node{Value: 4}
	ebt1.Right.Left = &Node{Value: 6}
	ebt1.Right.Right = &Node{Value: 5}
	ebt1.Left.Left.Left = &Node{Value: 1}

	ebt2 = &Node{Value: 3}
	ebt2.Left = &Node{Value: 5}
	ebt2.Right = &Node{Value: 7}
	ebt2.Left.Left = &Node{Value: 2}
	ebt2.Left.Right = &Node{Value: 4}
	ebt2.Right.Left = &Node{Value: 6}
	ebt2.Right.Right = &Node{Value: 8}
	ebt2.Left.Left.Left = &Node{Value: 1}

}

func TestFindTwoErrNode(t *testing.T) {
	wants := map[*Node][]int{
		ebt:  {6, 5},
		ebt1: {8, 5},
		ebt2: {5, 3},
	}
	for bt, want := range wants {
		t.Log(fmt.Sprintf("FindTwoErrNode  want: %v", want))
		PrintTree(bt)
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
		PrintTree(RecoverTree(bt))
		if fmt.Sprintf("%v", want) == fmt.Sprintf("%v", tmp) {
			t.Log("FindTwoErrNode is ok")
		} else {
			t.Error("FindTwoErrNode is not ok, result:", fmt.Sprintf("%v", tmp))
		}
	}

}
