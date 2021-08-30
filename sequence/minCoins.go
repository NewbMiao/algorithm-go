package sequence

import (
	"math"
)

// arr代表不重复的货币，可使用多次，求aim值得最小组合数.
func MinCoins1(arr []int, aim int) (c int) {
	if len(arr) == 0 || aim < 0 {
		return -1
	}
	n := len(arr)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, aim+1)
	}
	//!important 无对应组合设置为max，防止求最小值时，被0替换
	max := math.MaxInt32
	// dp[i][0]: aim为0时，对应组合都为0，无需处理

	// dp[0][j]: aim为 j=n*arr[i] 时，dp[0][j]=n  或者   dp[0][j-arr[0]] + 1
	for j := 1; j <= aim; j++ {
		dp[0][j] = max
		if j >= arr[0] && dp[0][j-arr[0]] != max { //  j%arr[0] == 0
			dp[0][j] = dp[0][j-arr[0]] + 1
		}
	}
	// dp[i][j]=min{dp[i-1][j],dp[i][j-arr[i]]+1}
	for i := 1; i < n; i++ {
		for j := 1; j <= aim; j++ {
			left := max
			if j >= arr[i] && dp[i][j-arr[i]] != max {
				left = dp[i][j-arr[i]] + 1
			}
			dp[i][j] = int(math.Min(float64(dp[i-1][j]), float64(left)))
		}
	}
	if dp[n-1][aim] == max {
		return -1
	}
	return dp[n-1][aim]
}

func MinCoins2(arr []int, aim int) (c int) {
	if len(arr) == 0 || aim < 0 {
		return -1
	}
	n := len(arr)
	// 压缩dp为行记录
	dp := make([]int, aim+1)
	//!important 无对应组合设置为max，防止求最小值时，被0替换
	max := math.MaxInt32
	// dp[0]: aim为0时，对应组合都为0，无需处理

	// dp[j]: aim为 j=n*arr[i] 时，dp[j]=n  或者   dp[j-arr[0]] + 1
	for j := 1; j <= aim; j++ {
		dp[j] = max
		if j >= arr[0] && dp[j-arr[0]] != max { //  j%arr[0] == 0
			dp[j] = dp[j-arr[0]] + 1
		}
	}
	// dp[j]=min{dp[j],dp[j-arr[i]]+1}
	for i := 1; i < n; i++ {
		for j := 1; j <= aim; j++ {
			left := max
			if j >= arr[i] && dp[j-arr[i]] != max {
				left = dp[j-arr[i]] + 1
			}
			dp[j] = int(math.Min(float64(dp[j]), float64(left)))
		}
	}
	if dp[aim] == max {
		return -1
	}
	return dp[aim]
}

//----------------------------------------------------------------

// arr代表可重复的货币，每张可使用1次，求aim值得最小组合数.
func MinCoins3(arr []int, aim int) (c int) {
	if len(arr) == 0 || aim < 0 {
		return -1
	}
	n := len(arr)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, aim+1)
	}
	//!important 无对应组合设置为max，防止求最小值时，被0替换
	max := math.MaxInt32
	// dp[i][0]: aim为0时，对应组合都为0，无需处理

	// dp[0][j]: aim为 arr[i] 时，dp[0][j]=1
	for j := 1; j <= aim; j++ {
		dp[0][j] = max
		if j == arr[0] {
			dp[0][j] = 1
		}
	}
	// dp[i][j]=min{dp[i-1][j],dp[i-1][j-arr[i]]+1}
	for i := 1; i < n; i++ {
		for j := 1; j <= aim; j++ {
			left := max
			if j >= arr[i] && dp[i][j-arr[i]] != max {
				left = dp[i-1][j-arr[i]] + 1
			}
			dp[i][j] = int(math.Min(float64(dp[i-1][j]), float64(left)))
		}
	}
	if dp[n-1][aim] == max {
		return -1
	}
	return dp[n-1][aim]
}

// arr代表可重复的货币，每张可使用1次，求aim值得最小组合数.
func MinCoins4(arr []int, aim int) (c int) {
	if len(arr) == 0 || aim < 0 {
		return -1
	}
	n := len(arr)
	dp := make([]int, aim+1)
	//!important 无对应组合设置为max，防止求最小值时，被0替换
	max := math.MaxInt32
	// dp[0]: aim为0时，对应组合都为0，无需处理

	// dp[j]: aim为 arr[i] 时，dp[j]=1
	for j := 1; j <= aim; j++ {
		dp[j] = max
		if j == arr[0] {
			dp[j] = 1
		}
	}
	// 按行滚动：dp[j]=min{dp[j],dp[j-arr[i]]+1}
	for i := 1; i < n; i++ {
		for j := 1; j <= aim; j++ {
			left := max
			if j >= arr[i] && dp[j-arr[i]] != max {
				left = dp[j-arr[i]] + 1
			}
			dp[j] = int(math.Min(float64(dp[j]), float64(left)))
		}
	}
	if dp[aim] == max {
		return -1
	}
	return dp[aim]
}
