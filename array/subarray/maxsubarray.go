package subarray

// 无序数组，求和为target的最大子数组
func MaxSubArraySum(d []int, target int) (res []int) {
	sumMap := map[int]int{}
	sumMap[0] = -1 // 设置和为0的索引位置

	var sum, preSum int
	for i, v := range d {
		sum += v
		sumMap[sum] = i
		preSum = sum - target
		if preSum < 0 {
			continue
		}

		// 若差值存在于sumMap，则返回（from+1,i）子数组
		if from, ok := sumMap[preSum]; ok {
			return d[from+1 : i+1]
		}
	}
	return
}
