package binary

//一组整型数据，这些数据中，其中有一个数只出现了一次，其他的数都出现了两次，找出出现一次的数.
func getNoRepeatNum(arr []int) (repeatNum int) {
	for _, v := range arr {
		repeatNum ^= v
	}
	return
}
