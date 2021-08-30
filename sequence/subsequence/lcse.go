package subsequence

import (
	"math"
)

// 最长公共子序列.
func GetLcse(str1, str2 string) string {
	if str1 == "" || str2 == "" {
		return ""
	}
	dp := getDpL(str1, str2)
	i := len(str1) - 1
	j := len(str2) - 1
	res := make([]uint8, dp[i][j])
	llen := dp[i][j] - 1
	for llen >= 0 {
		if i > 0 && dp[i][j] == dp[i-1][j] {
			i--
		} else if j > 0 && dp[i][j] == dp[i][j-1] {
			j--
		} else {
			res[llen] = str1[i]
			llen--
			i--
			j--
		}
	}
	return string(res)
}

func getDpL(str1, str2 string) (dp [][]int) {
	dp = make([][]int, len(str1))
	for k := range dp {
		dp[k] = make([]int, len(str2))
	}
	if str1[0] == str2[0] {
		dp[0][0] = 1
	}
	// 计算第一列
	for i := 1; i < len(str1); i++ {
		if str2[0] == str1[i] {
			dp[i][0] = 1
		}
		dp[i][0] = int(math.Max(float64(dp[i][0]), float64(dp[i-1][0])))
	}
	// 计算第一行
	for j := 1; j < len(str2); j++ {
		if str1[0] == str2[j] {
			dp[j][0] = 1
		}
		dp[0][j] = int(math.Max(float64(dp[0][j]), float64(dp[0][j-1])))
	}
	// 不相等则去前一行或前一列最大值，相等则取左上角+1
	for i := 1; i < len(str1); i++ {
		for j := 1; j < len(str2); j++ {
			dp[i][j] = int(math.Max(float64(dp[i-1][j]), float64(dp[i][j-1])))
			if str2[j] == str1[i] {
				dp[i][j] = int(math.Max(float64(dp[i][j]), float64(dp[i-1][j-1]+1)))
			}
		}
	}
	return dp
}
