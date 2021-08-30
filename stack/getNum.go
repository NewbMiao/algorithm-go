package stack

import "container/list"

func getNum(arr []int, num int) (res int) {
	l := len(arr)
	if l == 0 {
		return
	}
	// 双端队列维护当前arr[i,j]内的最大值和最小值
	qmax := list.New()
	qmin := list.New()
	var i, j int
	for i < l {
		for j < l {
			// j右移
			for qmax.Len() != 0 && arr[qmax.Back().Value.(int)] <= arr[j] {
				qmax.Remove(qmax.Back())
			}
			qmax.PushBack(j)

			for qmin.Len() != 0 && arr[qmin.Back().Value.(int)] >= arr[j] {
				qmin.Remove(qmin.Back())
			}
			qmin.PushBack(j)

			if arr[qmax.Front().Value.(int)]-arr[qmin.Front().Value.(int)] > num {
				break
			}
			j++
		}
		// i右移
		if qmax.Front().Value.(int) == i {
			qmax.Remove(qmax.Front())
		}
		if qmin.Front().Value.(int) == i {
			qmin.Remove(qmin.Front())
		}
		res += j - i
		i++
	}
	return res
}
