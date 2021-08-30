package sequence

import (
	"math"
)

//把str1变为str2需要的最小代价（dc：删除代价，rc：替换代价，ic：插入代价）
func MinEditCost(str1, str2 string, dc, rc, ic int) (c int) {
	row := len(str1) + 1
	col := len(str2) + 1

	dp := make([][]int, row)
	for k := range dp {
		dp[k] = make([]int, col)
	}
	//第一列，str1[i] => ''
	for i := 1; i < row; i++ {
		dp[i][0] = dc * i
	}
	//第一行，'' => str2[0]
	for j := 1; j < col; j++ {
		dp[0][j] = ic * j
	}
	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			if str1[i-1] == str2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = dp[i-1][j-1] + rc
			}
			//str1比str2长，删除一个字符
			dp[i][j] = int(math.Min(float64(dp[i-1][j]+dc), float64(dp[i][j])))
			//str1比str2短，插入一个字符
			dp[i][j] = int(math.Min(float64(dp[i][j-1]+ic), float64(dp[i][j])))
		}
	}

	return dp[row-1][col-1]
}

func MinEditCost2(str1, str2 string, dc, rc, ic int) (c int) {
	longs := str1
	shorts := str2
	if len(str1) < len(str2) {
		longs = str2
		shorts = str1
	}
	//str2长,作为行，则替换ic，dc
	if longs == str2 {
		ic, dc = dc, ic
	}
	dp := make([]int, len(shorts)+1)
	//第一行，'' => str2[0]
	for i := 1; i <= len(shorts); i++ {
		dp[i] = ic * i
	}

	for i := 1; i <= len(longs); i++ {
		pre := dp[0]
		dp[0] = dc * i
		for j := 1; j <= len(shorts); j++ {
			tmp := dp[j] //记录没更新前的dp[j]
			if longs[i-1] == shorts[j-1] {
				dp[j] = pre
			} else {
				dp[j] = pre + rc
			}
			//str1比str2长，删除一个字符
			dp[j] = int(math.Min(float64(tmp+dc), float64(dp[j])))
			//str1比str2短，插入一个字符
			dp[j] = int(math.Min(float64(dp[j-1]+ic), float64(dp[j])))
			pre = tmp //记录之前dp[j],下一次pre相当于dp[i-1][j-1]
		}
	}

	return dp[len(shorts)]
}
