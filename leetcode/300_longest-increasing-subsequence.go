package leetcode

import (
	"sort"
)

func lengthOfLIS(nums []int) int {
	l := len(nums)
	if l <= 1 {
		return l
	}

	dp := make([]int, l)
	dp[0] = 1
	res := 1
	for i := 1; i < l; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = Max(dp[j]+1, dp[i])
			}
		}
		res = Max(dp[i], res)
	}
	return res
}

func lengthOfLISGreedy(nums []int) int {
	l := len(nums)
	if l <= 1 {
		return l
	}
	res := 1
	dp := make([]int, l)
	dp[0] = nums[0]
	for i := 1; i < l; i++ {
		if nums[i] == dp[res-1] {
			continue
		}
		if nums[i] > dp[res-1] {
			dp[res] = nums[i]
			res++
		} else {
			k := sort.SearchInts(dp[:res], nums[i])
			dp[k] = nums[i]
		}
	}
	return res
}
