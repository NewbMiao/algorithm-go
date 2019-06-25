package traversal

import (
	"testing"
	. "btree"
)

var pbt *Node

func init() {
	pbt = &Node{Value: 1}
	pbt.Left = &Node{Value: 2}
	pbt.Right = &Node{Value: 3}
	pbt.Left.Left = &Node{Value: 4}
	pbt.Right.Left = &Node{Value: 5}
	pbt.Right.Right = &Node{Value: 6}
	pbt.Right.Left.Left = &Node{Value: 7}
	pbt.Right.Left.Right = &Node{Value: 8}

	PrintTree(pbt)

}

func TestPrintByLevel(t *testing.T) {
	PrintByLevel(pbt)
}

func TestPrintByZigZag(t *testing.T) {
	PrintByZigZag(pbt)
}