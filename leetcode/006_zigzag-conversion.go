package leetcode

func convert(s string, numRows int) string {
	if numRows <= 1 {
		return s
	}
	res := make([][]rune, numRows)
	i := 0
	flag := -1
	for _, v := range s {
		res[i] = append(res[i], v)
		if i == 0 || i == numRows-1 {
			flag = -flag
		}
		i += flag
	}
	output := ""
	for _, v := range res {
		output += string(v)
	}
	return output
}

func convertEven(s string, numRows int) string {
	if numRows <= 1 {
		return s
	}
	res := make([][]rune, numRows)
	var i int
	var isEven bool
	for k, v := range s {
		isEven = (k/(numRows-1))%2 == 0
		i = k % (numRows - 1)
		if isEven {
			res[i] = append(res[i], v)
		} else {
			res[numRows-1-i] = append(res[numRows-1-i], v)
		}
	}
	output := ""
	for _, v := range res {
		output += string(v)
	}
	return output
}
