package subtree

import (
	. "github.com/NewbMiao/AlgorithmCodeNote/kit/btree"
	"github.com/NewbMiao/AlgorithmCodeNote/kit/linkedlist"
	"math"
)

//非递归搜索
func GetBSTTopSize0(bt *Node) (size int) {
	if bt == nil {
		return
	}
	size = getNodeBSTTopSize(bt)
	size = int(math.Max(float64(GetBSTTopSize0(bt.Left)), float64(size)))
	size = int(math.Max(float64(GetBSTTopSize0(bt.Right)), float64(size)))
	return
}
func getNodeBSTTopSize(bt *Node) (size int) {
	if bt == nil {
		return
	}

	queue := linkedlist.NewList()
	queue.Push(bt)
	for !queue.IsEmpty() {
		tmp := queue.Pop()
		if tmp == nil {
			continue
		}
		node, ok := tmp.(*Node)
		if !ok || node == nil {
			continue
		}

		cur := bt
		//搜索该节点
		for cur != nil && cur != node {
			if node.Value < cur.Value {
				cur = cur.Left
			} else {
				cur = cur.Right
			}
		}
		if cur == node {
			size += 1
			queue.Push(node.Left)
			queue.Push(node.Right)
		}

	}
	return
}

//递归搜索
func GetBSTTopSize1(bt *Node) (size int) {
	if bt == nil {
		return
	}
	size = getNodeBSTTopSizeRecur(bt, bt)
	size = int(math.Max(float64(GetBSTTopSize1(bt.Left)), float64(size)))
	size = int(math.Max(float64(GetBSTTopSize1(bt.Right)), float64(size)))
	return
}
func getNodeBSTTopSizeRecur(bt, cur *Node) int {
	if isBSTNode(bt, cur) {
		return getNodeBSTTopSizeRecur(bt, cur.Left) + getNodeBSTTopSizeRecur(bt, cur.Right) + 1
	}
	return 0
}

func isBSTNode(bt, cur *Node) bool {
	if bt == nil || cur == nil {
		return false
	}
	if bt == cur {
		return true
	}
	if bt.Value > cur.Value {
		return isBSTNode(bt.Left, cur)
	} else {
		return isBSTNode(bt.Right, cur)
	}
}

type record struct {
	Lr int
	Rr int
}

//拓扑贡献记录
func GetBSTTopSize2(bt *Node) (size int) {
	if bt == nil {
		return
	}
	rMap := map[*Node]*record{}
	return posOrderRecordSize(bt, rMap)
}

func posOrderRecordSize(bt *Node, rMap map[*Node]*record) (size int) {
	if bt == nil {
		return
	}
	ls := posOrderRecordSize(bt.Left, rMap)
	rs := posOrderRecordSize(bt.Right, rMap)
	//修改上层拓扑记录
	modifySizeMap(bt.Left, bt.Value, rMap, true)
	modifySizeMap(bt.Right, bt.Value, rMap, false)
	lr := rMap[bt.Left]
	rr := rMap[bt.Right]
	curR := &record{}
	if lr != nil {
		curR.Lr = lr.Lr + lr.Rr + 1
	}
	if rr != nil {
		curR.Rr = rr.Lr + rr.Rr + 1
	}
	rMap[bt] = curR
	return int(math.Max(float64(curR.Lr+curR.Rr+1), math.Max(float64(ls), float64(rs))))
}

func modifySizeMap(n *Node, v int, rMap map[*Node]*record, isLeft bool) (minus int) {
	if n == nil {
		return
	}
	r := rMap[n]
	if r == nil {
		return
	}
	if (isLeft && n.Value > v) || (!isLeft && n.Value < v) {
		delete(rMap, n)
		return r.Lr + r.Rr + 1
	} else {
		//左子树左边界 && 右子树右边界无需修改贡献记录

		if isLeft { //左子树考察其右边界是否满足小于根节点
			minus = modifySizeMap(n.Right, v, rMap, isLeft)
			r.Rr -= minus
		} else { //右子树考察其左边界是否满足大于根节点
			minus = modifySizeMap(n.Left, v, rMap, isLeft)
			r.Lr -= minus
		}
		rMap[n] = r
		return
	}
}
