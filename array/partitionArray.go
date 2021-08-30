package array

// 将有序数组无重复部分移到左侧.
func leftUnique(arr []int) {
	l := len(arr)
	if l <= 1 {
		return
	}
	u := 0 // [0,u] 有序无重复区间
	i := 0 // [u+1，i] 有序区间

	for i != l {
		if arr[i] != arr[u] {
			arr[u+1], arr[i] = arr[i], arr[u+1] // 将u加入左无重复区间
			u++
		}
		i++
	}
}

func pSort(arr []int) {
	left := -1
	index := 0
	right := len(arr)

	for index < right {
		if arr[index] == 0 { // 左区域
			left++
			arr[index], arr[left] = arr[left], arr[index]
			index++
		} else if arr[index] == 2 { // 右区域
			right--
			arr[index], arr[right] = arr[right], arr[index]
			// 交换后不确定index位置大小，下一次判断时处理
		} else { // 中区域
			index++
		}
	}
}
