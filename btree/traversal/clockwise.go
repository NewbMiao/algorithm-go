package traversal

import (
	"github.com/NewbMiao/algorithm-go/kit/btree"
)

//逆时针打印树边界节点
func CounterClockWise(bt *btree.Node) (res []int) {
	height := btree.GetHeight(bt, 0)
	eMap := make(map[int][2]*btree.Node, height)
	btree.GetEdgeMap(bt, eMap, 0)
	//打印左边界
	for i := 0; i < height; i++ {
		res = append(res, eMap[i][0].Value)
	}
	//先序遍历叶子节点中非边界节点
	res = append(res, getLeafNotUnMap(bt, 0, eMap)...)
	//倒着打印右边界
	for i := height - 1; i > 0; i-- {
		if eMap[i][0].Value == eMap[i][1].Value {
			continue
		}
		res = append(res, eMap[i][1].Value)
	}
	return
}

//逆时针打印树边界节点（根节点，树左边界延伸，叶子节点，右边界延伸）
func CounterClockWise2(bt *btree.Node) (res []int) {
	if bt == nil {
		return
	}
	res = append(res, bt.Value)
	if bt.Left != nil && bt.Right != nil {
		res = append(res, getLeftEdge(bt.Left, true)...)
		res = append(res, getRightEdge(bt.Right, true)...)
	} else if bt.Left != nil {
		res = append(res, CounterClockWise2(bt.Left)...)
	} else {
		res = append(res, CounterClockWise2(bt.Right)...)
	}
	return
}

func getLeftEdge(bt *btree.Node, print bool) (res []int) {
	if bt == nil {
		return
	}
	if print || (bt.Left == nil && bt.Right == nil) {
		res = append(res, bt.Value)
	}
	res = append(res, getLeftEdge(bt.Left, print)...)
	rPrint := false
	if print && bt.Left == nil { //若打印时无左节点则打印右节点
		rPrint = true
	}
	res = append(res, getLeftEdge(bt.Right, rPrint)...)
	return
}

func getRightEdge(bt *btree.Node, print bool) (res []int) {
	if bt == nil {
		return
	}
	lPrint := false
	if print && bt.Right == nil { //若打印时无右节点则打印左节点
		lPrint = true
	}
	res = append(res, getRightEdge(bt.Left, lPrint)...)
	res = append(res, getRightEdge(bt.Right, print)...)

	if print || (bt.Left == nil && bt.Right == nil) {
		res = append(res, bt.Value)
	}

	return
}

func getLeafNotUnMap(bt *btree.Node, l int, eMap map[int][2]*btree.Node) (res []int) {
	if bt == nil {
		return
	}

	if bt.Right == nil && bt.Left == nil &&
		bt != eMap[l][0] && bt != eMap[l][1] {
		res = append(res, bt.Value)
	}
	res = append(res, getLeafNotUnMap(bt.Left, l+1, eMap)...)
	res = append(res, getLeafNotUnMap(bt.Right, l+1, eMap)...)
	return
}
