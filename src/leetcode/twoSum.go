package leetcode

//https://leetcode-cn.com/problems/two-sum/submissions/
/*
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。
*/
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
