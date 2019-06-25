package btree

import (
	"math"
)

//是否为平衡二叉树，左右子树高差不超过1
func IsBalanceTree(bt *Node) (r bool) {
	res := new(bool)
	*res = true
	getSubtreeHeight(bt, 1, res)
	return *res
}

func getSubtreeHeight(bt *Node, level int, res *bool) (h int) {
	if bt == nil {
		return level
	}
	lh := getSubtreeHeight(bt.Left, level+1, res)
	if !*res {
		return level
	}
	rh := getSubtreeHeight(bt.Right, level+1, res)
	if !*res {
		return level
	}
	if math.Abs(float64(lh-rh)) > 1 {
		*res = false
	}
	return int(math.Max(float64(lh), float64(rh)))
}
