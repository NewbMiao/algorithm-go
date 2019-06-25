package subtree

import (
	. "btree"
	"math"
)

func GetMaxSubBST(bt *Node) *Node {
	record := map[int]int{} // size min max
	//record := &[3]int{} // size min max 若为数组需传递指针

	return posOrderRecord(bt, record)

}

func posOrderRecord(bt *Node, record map[int]int) *Node {

	if bt == nil {
		record[0] = 0
		record[1] = math.MaxInt64
		record[2] = math.MinInt64
		return nil
	}
	value, left, right := bt.Value, bt.Left, bt.Right
	lBST := posOrderRecord(left, record)
	lSize, lMin, lMax := record[0], record[1], record[2]
	rBST := posOrderRecord(right, record)
	rSize, rMin, rMax := record[0], record[1], record[2]

	record[1] = int(math.Min(float64(lMin), float64(value)))
	record[2] = int(math.Max(float64(rMax), float64(value)))
	if left == lBST && right == rBST && lMax < value && value < rMin {
		record[0] = lSize + rSize + 1
		return bt
	}
	record[0] = int(math.Max(float64(lSize), float64(rSize)))
	if lSize > rSize {
		return lBST
	}
	return rBST
}
