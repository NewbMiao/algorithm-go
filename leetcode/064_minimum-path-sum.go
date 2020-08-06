package leetcode

// optimized version see: sequence/minPathSum.go
func minPathSum(grid [][]int) int {
	m := len(grid)
	if m < 1 {
		return 0
	}
	n := len(grid[0])
	dp := make([][]int, m)
	for k := range dp {
		dp[k] = make([]int, n)
	}

	dp[0][0] = grid[0][0]

	for i := 1; i < n; i++ {
		dp[0][i] = dp[0][i-1] + grid[0][i]
	}

	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
		for j := 1; j < n; j++ {
			dp[i][j] = Min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[m-1][n-1]
}
