package sequence

// 母牛永远不死，母牛每年生一头小母牛。第一年1头母牛，第二年母牛开始生小母牛。
// 每只小母牛3年后成熟又可以生小母牛。给定整数N，求N年后牛数量

//c(n)=c(n-1)+c(n-3)
func Cow1(n int) int {
	if n < 1 {
		return 0
	}
	if n <= 3 {
		return n
	}
	return Cow1(n-1) + Cow1(n-3)
}

func Cow2(n int) int {
	if n < 1 {
		return 0
	}
	if n <= 3 {
		return n
	}
	prepre := 1
	pre := 2
	res := 3
	tmp1 := 0
	tmp2 := 0
	for i := 4; i <= n; i++ {
		tmp1 = res
		tmp2 = pre
		res += prepre
		pre = tmp1
		prepre = tmp2
	}
	return res
}

func Cow3(n int) int {
	if n < 1 {
		return 0
	}
	if n <= 3 {
		return n
	}
	base := [][]int{{1, 1, 0}, {0, 0, 1}, {1, 0, 0}}
	res := matrixPower(base, n-3)
	return 3*res[0][0] + 2*res[1][0] + res[2][0]
}
