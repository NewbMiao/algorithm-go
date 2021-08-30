package sequence

import (
	"math"
)

// 矩阵中向右或向下走，到达右下角的最短路径和.
func MinPathSum1(m [][]int) int {
	if len(m) == 0 {
		return 0
	}
	dp := make([][]int, len(m))
	for i := range dp {
		dp[i] = make([]int, len(m[0]))
	}
	var i, j int
	for i = range m {
		for j = range m[i] {
			v := m[i][j]
			if i == 0 && j == 0 {
				dp[i][j] = v
				continue
			}
			// 第一行
			if i == 0 && j > 0 {
				dp[i][j] += dp[i][j-1] + v
				continue
			}
			// 第一列
			if j == 0 && i > 0 {
				dp[i][j] += dp[i-1][j] + v
				continue
			}
			// 其他路径，经过左边和上边值
			dp[i][j] = int(math.Min(float64(dp[i-1][j]), float64(dp[i][j-1]))) + v
		}
	}
	return dp[len(m)-1][len(m[0])-1]
}

func MinPathSum2(m [][]int) int {
	if len(m) == 0 {
		return 0
	}
	more := int(math.Max(float64(len(m)), float64(len(m[0]))))
	less := int(math.Min(float64(len(m)), float64(len(m[0]))))
	dp := make([]int, less)
	rowMore := true
	if more == len(m[0]) {
		rowMore = false
	}

	dp[0] = m[0][0]
	// 初始化第一行/列
	for i := 1; i < less; i++ {
		if rowMore {
			dp[i] = dp[i-1] + m[0][i]
		} else {
			dp[i] = dp[i-1] + m[i][0]
		}
	}

	// 按行或列滚动计算
	for i := 1; i < more; i++ {
		// 计算第一次滚动
		if rowMore {
			dp[0] += m[i][0]
		} else {
			dp[0] += m[0][i]
		}
		// 最小值必与前两次滚动最小值有关
		for j := 1; j < less; j++ {
			if rowMore {
				dp[j] = int(math.Min(float64(dp[j-1]), float64(dp[j]))) + m[i][j]
			} else {
				dp[j] = int(math.Min(float64(dp[j-1]), float64(dp[j]))) + m[j][i]
			}
		}
	}

	return dp[less-1]
}
