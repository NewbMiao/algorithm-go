package leetcode

import (
	"fmt"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	ds := []string{
		"23",
	}
	wants := [][]string{
		{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
	}
	for k, v := range ds {
		t.Log(fmt.Sprintf("letterCombinations, input arr %v, want: %v", v, wants[k]))
		r := letterCombinations(v)
		if fmt.Sprint(r) == fmt.Sprint(wants[k]) {
			t.Log("letterCombinations is ok")
		} else {
			t.Error("letterCombinations is not ok, result:", fmt.Sprintf("%v", r))
		}

		r = letterCombinations1(v)
		if fmt.Sprint(r) == fmt.Sprint(wants[k]) {
			t.Log("letterCombinations1 is ok")
		} else {
			t.Error("letterCombinations1 is not ok, result:", fmt.Sprintf("%v", r))
		}

		r = letterCombinations2(v)
		if fmt.Sprint(r) == fmt.Sprint(wants[k]) {
			t.Log("letterCombinations2 is ok")
		} else {
			t.Error("letterCombinations2 is not ok, result:", fmt.Sprintf("%v", r))
		}
	}
}
