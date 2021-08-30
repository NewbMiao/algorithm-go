package stack

import (
	"container/list"
)

// 获取滑动窗口中最大值数组
// 使用内置双链表，remove可优化为移除队首和移除队尾.
func getMaxWindow(arr []int, w int) (res []int) {
	if len(arr) == 0 || w < 1 || len(arr) < w {
		return
	}
	pmax := list.New()
	res = make([]int, 0)

	for i := range arr {
		// 维护大值到队首（移除队尾）
		for pmax.Len() != 0 && arr[pmax.Back().Value.(int)] <= arr[i] {
			pmax.Remove(pmax.Back())
		}
		pmax.PushBack(i)

		// 滑动窗口，删除头部过期索引（移除队首）
		if i-w == pmax.Front().Value.(int) {
			pmax.Remove(pmax.Front())
		}

		// 记录当前窗口最大值
		if i >= w-1 {
			res = append(res, arr[pmax.Front().Value.(int)])
		}
	}

	return
}
