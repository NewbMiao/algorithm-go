package sequence

//arr中为正数且不重复，每个值代表一种货币，可使用任意张，求值为aim的换钱有多少种
func CoinsWay1(arr []int, aim int) int {
	if len(arr) == 0 || aim <= 0 {
		return 0
	}
	return process1(arr, 0, aim)
}

func process1(arr []int, index, aim int) (res int) {
	res = 0
	if index == len(arr) { //最后一次
		if aim == 0 {
			res = 1
		}
	} else {
		for i := 0; arr[index]*i <= aim; i++ {
			//计算剩余钱数是否能组合
			res += process1(arr, index+1, aim-arr[index]*i)
		}
	}

	return
}

func CoinsWay2(arr []int, aim int) int {
	if len(arr) == 0 || aim <= 0 {
		return 0
	}
	rMap := make([][]int, len(arr)+1) //0未记录；-1无对应组合
	for i := range rMap {
		rMap[i] = make([]int, aim+1)
	}
	return process2(arr, &rMap, 0, aim)
}

func process2(arr []int, rMap *[][]int, index, aim int) (res int) {
	res = 0
	if index == len(arr) {
		if aim == 0 {
			res = 1
		}
	} else {
		for i := 0; arr[index]*i <= aim; i++ {
			rv := (*rMap)[index+1][aim-arr[index]*i]
			//计算剩余钱数是否能组合
			if rv == 0 {
				res += process2(arr, rMap, index+1, aim-arr[index]*i)
			} else {
				if rv != -1 {
					res += rv
				}
			}
		}
	}
	if res == 0 {
		(*rMap)[index][aim] = -1
	} else {
		(*rMap)[index][aim] = res
	}
	return
}

func CoinsWay3(arr []int, aim int) int {
	if len(arr) == 0 || aim <= 0 {
		return 0
	}
	n := len(arr)
	dp := make([][]int, len(arr))
	for i := range dp {
		dp[i] = make([]int, aim+1)
	}
	//aim=0 所有钱组合为1
	for i := 0; i < n; i++ {
		dp[i][0] = 1
	}

	//arr[0]的组合为 0*c0 1*c0 ...
	for i := 1; i*arr[0] <= aim; i++ {
		dp[0][i*arr[0]] = 1
	}

	//
	for i := 1; i < n; i++ {
		for j := 1; j <= aim; j++ {
			num := 0
			for k := 0; j-arr[i]*k >= 0; k++ {
				num += dp[i-1][j-arr[i]*k]
			}
			dp[i][j] = num
		}
	}

	return dp[n-1][aim]
}

func CoinsWay4(arr []int, aim int) int {
	if len(arr) == 0 || aim <= 0 {
		return 0
	}
	n := len(arr)
	dp := make([][]int, len(arr))
	for i := range dp {
		dp[i] = make([]int, aim+1)
	}
	//aim=0 所有钱组合为1
	for i := 0; i < n; i++ {
		dp[i][0] = 1
	}

	//arr[0]的组合为 0*c0 1*c0 ...
	for i := 1; i*arr[0] <= aim; i++ {
		dp[0][i*arr[0]] = 1
	}

	//
	for i := 1; i < n; i++ {
		for j := 1; j <= aim; j++ {
			dp[i][j] = dp[i-1][j]
			if j-arr[i] >= 0 {
				dp[i][j] += dp[i][j-arr[i]]
			}
		}
	}

	return dp[n-1][aim]
}

func CoinsWay5(arr []int, aim int) int {
	if len(arr) == 0 || aim <= 0 {
		return 0
	}
	n := len(arr)
	dp := make([]int, aim+1)
	//aim=0 所有钱组合为1
	//arr[0]的组合为 0*c0 1*c0 ...
	for i := 0; i*arr[0] <= aim; i++ {
		dp[i*arr[0]] = 1
	}

	//
	for i := 1; i < n; i++ {
		for j := 1; j <= aim; j++ {
			if j-arr[i] >= 0 {
				dp[j] += dp[j-arr[i]]
			}
		}
	}

	return dp[aim]
}
