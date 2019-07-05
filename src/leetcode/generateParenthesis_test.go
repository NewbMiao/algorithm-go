package leetcode

import (
	"fmt"
	"testing"
)

func TestGenerateParenthesis(t *testing.T) {
	ds := []int{
		3,
	}
	ws := [][]string{
		{"((()))",
			"(()())",
			"(())()",
			"()(())",
			"()()()",},
	}
	for k, v := range ds {
		t.Log(fmt.Sprintf("GenerateParenthesis, input num %v, want: %v", v, ws[k]))
		r := generateParenthesis(v)
		if fmt.Sprint(r) == fmt.Sprint(ws[k]) {
			t.Log("GenerateParenthesis is ok")
		} else {
			t.Error("GenerateParenthesis is not ok, result:", fmt.Sprintf("%v", r))
		}

		r = generateParenthesis1(v)
		if fmt.Sprint(r) == fmt.Sprint(ws[k]) {
			t.Log("GenerateParenthesis1 is ok")
		} else {
			t.Error("GenerateParenthesis1 is not ok, result:", fmt.Sprintf("%v", r))
		}

	}
}
