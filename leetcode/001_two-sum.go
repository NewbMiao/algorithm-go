package leetcode

func twoSum(nums []int, target int) []int {
	if len(nums) == 0 {
		return nil
	}
	sumMap := map[int]int{}
	for k, v := range nums {
		if k1, ok := sumMap[target-v]; ok {
			return []int{k1, k}
		} else {
			sumMap[v] = k
		}
	}
	return nil
}
