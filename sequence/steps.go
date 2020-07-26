package sequence

//给定整数N，代表台阶数，可以一次跨2个或1个台阶，返回多少种走法

//s(n)=s(n-1)+s(n-2)
func Steps1(n int) int {
	if n < 1 {
		return 0
	}
	if n <= 2 {
		return n
	}
	return Steps1(n-1) + Steps1(n-2)
}

func Steps2(n int) int {
	if n < 1 {
		return 0
	}
	if n <= 2 {
		return n
	}
	tmp := 0
	pre := 1
	res := 2
	for i := 3; i <= n; i++ {
		tmp = res
		res += pre
		pre = tmp
	}

	return res
}

//矩阵
func Steps3(n int) int {
	base := [][]int{{1, 1}, {1, 0}}
	res := matrixPower(base, n-2)
	return 2*res[0][0] + res[1][0]
}
