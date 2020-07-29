package leetcode

func maxSubArray(nums []int) int {
	l := len(nums)
	if l < 1 {
		return 0
	}
	dp := make([]int, l)
	dp[0] = nums[0]
	result := dp[0]
	for i := 1; i < l; i++ {
		dp[i] = Max(dp[i-1]+nums[i], nums[i])
		result = Max(dp[i], result)
	}
	return result
}
