package leetcode

func minimumTotal(triangle [][]int) int {
	l := len(triangle)
	if l < 1 {
		return 0
	}
	if l == 1 {
		return triangle[0][0]
	}
	for i := 1; i < l; i++ {
		l2 := len(triangle[i])
		for j := 0; j < l2; j++ {
			if j == 0 {
				triangle[i][j] = triangle[i-1][0] + triangle[i][j]
			} else if j == l2-1 {
				triangle[i][j] = triangle[i-1][l2-2] + triangle[i][j]
			} else {
				triangle[i][j] = Min(triangle[i-1][j], triangle[i-1][j-1]) + triangle[i][j]
			}
		}
	}
	total := 1<<31 - 1
	for _, v := range triangle[l-1] {
		if total > v {
			total = v
		}
	}
	return total
}
