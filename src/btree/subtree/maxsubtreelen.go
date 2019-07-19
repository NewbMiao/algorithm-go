package subtree

import (
	. "kit/btree"
	"math"
)

func GetMaxSubtreeLen(bt *Node, sum int) (maxLen int) {
	sumMap := map[int]int{}
	sumMap[0] = 0 //important
	return getMaxLenPreOrder(bt, sum, 0, 1, 0, sumMap)
}

func getMaxLenPreOrder(bt *Node, sum, preSum, level, maxLen int, sumMap map[int]int) int {

	if bt == nil {
		return maxLen
	}
	curSum := preSum + bt.Value

	if _, ok := sumMap[curSum]; !ok {
		sumMap[curSum] = level
	}

	if preLevel, ok := sumMap[curSum-sum]; ok {
		maxLen = int(math.Max(float64(maxLen), float64(level-preLevel)))

	}
	maxLen = getMaxLenPreOrder(bt.Left, sum, curSum, level+1, maxLen, sumMap)
	maxLen = getMaxLenPreOrder(bt.Right, sum, curSum, level+1, maxLen, sumMap)

	//if v, ok := sumMap[curSum]; ok && v == level {
	//	delete(sumMap, curSum)
	//}

	return maxLen
}
