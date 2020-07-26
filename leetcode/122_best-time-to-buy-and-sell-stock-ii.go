package leetcode

func maxProfit(prices []int) int {
	num := len(prices)
	if num < 2 {
		return 0
	}
	max := 0
	for i := 0; i+1 < num; i++ {
		if prices[i+1] > prices[i] {
			max += prices[i+1] - prices[i]
		}
	}
	return max
}

func maxprofitDp(prices []int) int {
	num := len(prices)
	if num < 2 {
		return 0
	}

	dp := make([][2]int, num)
	dp[0][0] = 0          //sell out
	dp[0][1] = -prices[0] //buy in
	for i := 1; i < num; i++ {
		dp[i][0] = Max(dp[i-1][0], dp[i-1][1]+prices[i]) //sell out
		dp[i][1] = Max(dp[i-1][1], dp[i-1][0]-prices[i]) //buy in
	}
	return dp[num-1][0]
}
