package sequence

import (
	"math"
)

//龙与地下城游戏：左上角出发，只能向下或向右，每个值代表加血或减血，求到右下角至少剩1血量，骑士所需的初始血量
func DungeonGame(m [][]int) (b int) {
	if len(m) == 0 || len(m[0]) == 0 {
		return 1
	}
	row := len(m)
	col := len(m[0])
	dp := make([][]int, row)
	for k := range dp {
		dp[k] = make([]int, col)
	}

	//从右下角开始计算
	dp[row-1][col-1] = 1
	if m[row-1][col-1] < 0 {
		dp[row-1][col-1] = 1 - m[row-1][col-1]
	}

	//最后一行，从右向左
	for j := col - 2; j >= 0; j-- {
		dp[row-1][j] = int(math.Max(float64(dp[row-1][j+1]-m[row-1][j]), 1))
	}
	//从下到上
	for i := row - 2; i >= 0; i-- {
		dp[i][col-1] = int(math.Max(float64(dp[i+1][col-1]-m[i][col-1]), 1))
		for j := col - 2; j >= 0; j-- {
			down := math.Max(float64(dp[i+1][j]-m[i][j]), 1)
			right := math.Max(float64(dp[i][j+1]-m[i][j]), 1)
			dp[i][j] = int(math.Min(right, down))
		}
	}
	return dp[0][0]
}

func DungeonGame2(m [][]int) (b int) {
	if len(m) == 0 || len(m[0]) == 0 {
		return 1
	}
	row := len(m)
	col := len(m[0])
	more := row
	less := col
	rowMore := true
	if row < col {
		more = col
		less = row
		rowMore = false
	}
	dp := make([]int, less)

	//从右下角开始计算
	dp[less-1] = 1
	if m[row-1][col-1] < 0 {
		dp[less-1] = 1 - m[row-1][col-1]
	}
	//最后一行，从右向左
	for j := less - 2; j >= 0; j-- {
		if rowMore {
			dp[j] = int(math.Max(float64(dp[j+1]-m[more-1][j]), 1))
		} else {
			dp[j] = int(math.Max(float64(dp[j+1]-m[j][more-1]), 1))
		}
	}

	//从下到上
	for i := more - 2; i >= 0; i-- {
		row := i
		col := less - 1

		if !rowMore {
			row = less - 1
			col = i
		}
		dp[less-1] = int(math.Max(float64(dp[less-1]-m[row][col]), 1))

		for j := less - 2; j >= 0; j-- {
			row := i
			col := j
			if !rowMore {
				row = j
				col = i
			}
			down := math.Max(float64(dp[j]-m[row][col]), 1)
			right := math.Max(float64(dp[j+1]-m[row][col]), 1)
			dp[j] = int(math.Min(right, down))

		}
	}

	return dp[0]
}
