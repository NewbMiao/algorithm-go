package leetcode

import "sort"

var pMap = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

func letterCombinations(digits string) []string {
	if digits == "" {
		return nil
	}
	return backtrace("", digits)
}

func backtrace(letters, next_digit string) (res []string) {
	if next_digit == "" {
		res = append(res, letters)
		return
	}
	d := next_digit[:1]
	for _, v := range pMap[string(d)] {
		res = append(res, backtrace(letters+string(v), next_digit[1:])...)

	}
	return
}

//rune
func letterCombinations1(digits string) []string {
	if digits == "" {
		return nil
	}
	res := [][]rune{[]rune(""),}
	for _, d := range digits {
		next_res := [][]rune{}
		for _, letter := range pMap[string(d)] {
			for _, v := range res {
				//tmp := make([]rune, len(v))
				//copy(tmp, v) //拷贝rune，避免重复修改 v
				next_res = append(next_res, append([]rune(string(v)), letter))

			}

		}
		res = next_res
	}
	out := make([]string, len(res))
	for k := range res {
		out[k] = string(res[k])
	}
	sort.Strings(out)
	return out

}

//string
func letterCombinations2(digits string) []string {
	if digits == "" {
		return nil
	}
	res := []string{""}
	for _, d := range digits {
		next_res := []string{}
		for _, letter := range pMap[string(d)] {
			for _, v := range res {
				next_res = append(next_res, v+string(letter))
			}

		}
		res = next_res

	}

	sort.Strings(res)
	return res

}
