package subsequence

import (
	"math"
)

// 最长递增子序列
func GetLIS1(arr []int) (res []int) {
	if len(arr) == 0 {
		return
	}
	// 到数组arr位置i，之前最长递增子序列长度为n, 对应res的index为n-1
	dp := getDp1(arr)
	return genLIS(arr, dp)
}

func getDp1(arr []int) (dp []int) {
	dp = make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		dp[i] = 1
		for j := 1; j < i; j++ {
			if arr[j] < arr[i] {
				dp[i] = int(math.Max(float64(dp[i]), float64(dp[j]+1)))
			}
		}
	}
	return
}

func genLIS(arr []int, dp []int) (res []int) {
	maxK := 0
	for k, v := range dp {
		if v > dp[maxK] {
			maxK = k
		}
	}
	res = make([]int, dp[maxK])
	res[dp[maxK]-1] = arr[maxK]
	for j := maxK - 1; j >= 0; j-- {
		if arr[j] < arr[j+1] && dp[j+1]-dp[j] == 1 {
			res[dp[j]-1] = arr[j]
		}
	}
	return
}

func GetLIS2(arr []int) (res []int) {
	if len(arr) == 0 {
		return
	}
	return genLIS(arr, getDp2(arr))
}
func getDp2(arr []int) (dp []int) {
	n := len(arr)
	dp = make([]int, n)
	ends := make([]int, n) // 有效区：对应i位置数结尾的递增子序列
	right := 0             // 有效区右边界
	dp[0] = 1
	ends[0] = arr[0]
	var l, r int
	m := 0
	for i := 1; i < n; i++ {
		//二分查找不大于ends的位置
		l = 0
		r = right
		for l <= r {
			m = (l + r) / 2
			if arr[i] > ends[m] {
				l = m + 1
			} else {
				r = m - 1
			}
		}
		right = int(math.Max(float64(right), float64(l)))
		dp[i] = l + 1
		ends[l] = arr[i]
	}

	return
}
