package sequence

func Fibonacci1(n int) int {
	if n < 1 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	return Fibonacci1(n-1) + Fibonacci1(n-2)
}

func Fibonacci2(n int) int {
	if n < 1 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	res := 1
	pre := 1
	var tmp int
	for i := 3; i <= n; i++ {
		tmp = res
		res += pre
		pre = tmp
	}
	return res
}

// 斐波那契数列 f(n)=f(n-1)+f(n-2).
func Fibonacci3(n int) int {
	if n < 1 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	base := [][]int{{1, 1}, {1, 0}}
	res := matrixPower(base, n-2)
	return res[0][0] + res[1][0]
}

func matrixPower(m [][]int, p int) (res [][]int) {
	if len(m) == 0 {
		return
	}
	res = make([][]int, len(m))
	// 单位矩阵 [[1 0] [0 1]]
	for i := 0; i < len(m); i++ {
		res[i] = make([]int, len(m[0]))
		res[i][i] = 1
	}
	tmp := m
	for ; p != 0; p >>= 1 {
		if (p & 1) != 0 {
			res = multiMatrix(res, tmp)
		}
		tmp = multiMatrix(tmp, tmp)
	}
	return
}

func multiMatrix(m1, m2 [][]int) (res [][]int) {
	res = make([][]int, len(m1))
	for i := 0; i < len(m1); i++ {
		res[i] = make([]int, len(m2[0]))
	}

	for i := 0; i < len(m1); i++ {
		for j := 0; j < len(m2[0]); j++ {
			for k := 0; k < len(m2); k++ {
				res[i][j] += m1[i][k] * m2[k][j]
			}
		}
	}
	return
}
