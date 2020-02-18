package sequence

//判定aim是否顺序交错包含str1和str2
func IsStrCross(str1, str2, aim string) bool {
	if len(aim) != len(str1)+len(str2) {
		return false
	}
	dp := make([][]bool, len(str1)+1)
	for k := range dp {
		dp[k] = make([]bool, len(str2)+1)
	}

	dp[0][0] = true
	//第一列，无str1时，计算str2和aim的交错
	for i := 1; i <= len(str1); i++ {
		//顺序包含，所以当不相等时无需比较后续str
		if str1[i-1] != aim[i-1] {
			break
		}
		dp[i][0] = true

	}
	//第一行，无str2时，计算str1和aim的交错
	for j := 1; j <= len(str2); j++ {
		if str2[j-1] != aim[j-1] {
			break
		}
		dp[0][j] = true
	}
	for i := 1; i <= len(str1); i++ {
		for j := 1; j <= len(str2); j++ {
			if (dp[i-1][j] && str1[i-1] == aim[i+j-1]) ||
				(dp[i][j-1] && str2[j-1] == aim[i+j-1]) {
				dp[i][j] = true
			}
		}
	}
	return dp[len(str1)][len(str2)]
}

func IsStrCross2(str1, str2, aim string) bool {
	if len(aim) != len(str1)+len(str2) {
		return false
	}
	longs := str1
	shorts := str2
	if len(str1) < len(str2) {
		longs = str2
		shorts = str1
	}
	dp := make([]bool, len(shorts)+1)

	//第一行或第一列
	for i := 1; i <= len(shorts); i++ {
		//顺序包含，所以当不相等时无需比较后续str
		if str1[i-1] != aim[i-1] {
			break
		}
		dp[i] = true

	}
	for i := 1; i <= len(longs); i++ {
		dp[0] = dp[0] && longs[i-1] == aim[i-1]
		for j := 1; j <= len(shorts); j++ {
			if (dp[j] && longs[i-1] == aim[i+j-1]) ||
				(dp[j-1] && shorts[j-1] == aim[i+j-1]) {
				dp[j] = true
			}
		}
	}
	return dp[len(shorts)]
}
