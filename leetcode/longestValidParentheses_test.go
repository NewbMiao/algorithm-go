package leetcode

import (
	"fmt"
	"testing"
)

func TestLongestValidParentheses(t *testing.T) {
	ds := []string{
		"()",
		"(()",
		")()())",
		"()(()",
	}
	ws := []int{
		2,
		2,
		4,
		2,
	}

	for k, v := range ds {
		t.Log(fmt.Sprintf("longestValidParentheses, input str %v, want: %v", v, ws[k]))
		r := longestValidParentheses(v)
		if r == ws[k] {
			t.Log("longestValidParentheses is ok")
		} else {
			t.Error("longestValidParentheses is not ok, result:", fmt.Sprintf("%v", r))
		}

		r = longestValidParentheses1(v)
		if r == ws[k] {
			t.Log("longestValidParentheses1 is ok")
		} else {
			t.Error("longestValidParentheses1 is not ok, result:", fmt.Sprintf("%v", r))
		}

		r = longestValidParentheses2(v)
		if r == ws[k] {
			t.Log("longestValidParentheses2 is ok")
		} else {
			t.Error("longestValidParentheses2 is not ok, result:", fmt.Sprintf("%v", r))
		}
	}
}
