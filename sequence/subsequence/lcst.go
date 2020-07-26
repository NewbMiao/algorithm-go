package subsequence

//最长公共子串
func GetLcst(str1, str2 string) (res string) {
	if str1 == "" || str2 == "" {
		return
	}
	dp := GetDpLcst(str1, str2)
	mlen := 0
	endIndex := 0
	for i := range dp {
		for j := range dp[i] {
			if dp[i][j] > mlen {
				mlen = dp[i][j]
				endIndex = i
			}
		}
	}
	if mlen > 0 {
		res = str1[endIndex-mlen+1 : endIndex+1]
	}
	return
}

func GetDpLcst(str1, str2 string) (dp [][]int) {
	dp = make([][]int, len(str1))
	for k := range dp {
		dp[k] = make([]int, len(str2))
	}
	if str1[0] == str2[0] {
		dp[0][0] = 1
	}
	for i := 1; i < len(str1); i++ {
		if str2[0] == str1[i] {
			dp[i][0] = 1
		}
	}
	for j := 1; j < len(str2); j++ {
		if str1[0] == str2[j] {
			dp[j][0] = 1
		}
	}

	for i := 1; i < len(str1); i++ {
		for j := 1; j < len(str2); j++ {
			if str1[i] == str2[j] {
				dp[i][j] = dp[i-1][j-1] + 1
			}
		}
	}
	return
}

//按斜线，从右上到左下遍历
func GetLcst2(str1, str2 string) (res string) {
	if str1 == "" || str2 == "" {
		return
	}
	max := 0 //最大长度
	endIndex := 0
	row := 0
	col := len(str2) - 1
	for row < len(str1) {
		i := row
		j := col
		mlen := 0
		for i < len(str1) && j < len(str2) {
			if str2[j] == str1[i] {
				mlen++
			} else {
				mlen = 0
			}
			if mlen > max {
				max = mlen
				endIndex = i
			}
			i++
			j++
		}

		if col > 0 { //右向左
			col--
		} else { //上向下
			row++
		}
	}
	if max > 0 {
		res = str1[endIndex-max+1 : endIndex+1]
	}
	return res
}
