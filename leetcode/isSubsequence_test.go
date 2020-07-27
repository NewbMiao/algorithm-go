package leetcode

import (
	"fmt"
	"testing"
)

func TestIsSubsequence(t *testing.T) {
	d := [][]string{
		{"abc", "ahbgdc"},
		{"twn", "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxtxxxxxxxxxxxxxxxxxxxxwxxxxxxxxxxxxxxxxxxxxxxxxxn"},
		{"axc", "ahbgdc"},
	}
	wants := []bool{
		true,
		true,
		false,
	}
	for k, v := range d {
		t.Log(fmt.Sprintf("check s isSubsequence of t, input s %v, t %v, want: %v", v[0], v[1], wants[k]))
		r := isSubsequence(v[0], v[1])
		if r == wants[k] {
			t.Log("isSubsequence is ok")
		} else {
			t.Error("isSubsequence is not ok, result:", fmt.Sprintf("%v", r))
		}

		r = isSubsequence2(v[0], v[1])
		if r == wants[k] {
			t.Log("isSubsequence2 is ok")
		} else {
			t.Error("isSubsequence2 is not ok, result:", fmt.Sprintf("%v", r))
		}
		r = isSubsequence3(v[0], v[1])
		if r == wants[k] {
			t.Log("isSubsequence3 is ok")
		} else {
			t.Error("isSubsequence3 is not ok, result:", fmt.Sprintf("%v", r))
		}
	}
}
