package leetcode

import (
	"testing"
	"fmt"
)

func TestIsMatch(t *testing.T) {
	d := [][2]string{
		{"aa", "a"},
		{"aa", "a*"},
		{"ab", ".*"},
		{"aab", "c*a*b*"},
		{"mississippi", "mis*is*p*."},
	}
	wants := []bool{
		false,
		true,
		true,
		true,
		false,
	}
	for k, v := range d {
		t.Log(fmt.Sprintf("isMatch, input s %v, p %v, want: %v", v[0], v[1], wants[k]))
		r := isMatch(v[0], v[1])
		if r == wants[k] {
			t.Log("isMatch is ok")
		} else {
			t.Error("isMatch is not ok, result:", fmt.Sprintf("%v", r))
		}

		r = isMatch2(v[0], v[1])
		if r == wants[k] {
			t.Log("isMatch2 is ok")
		} else {
			t.Error("isMatch2 is not ok, result:", fmt.Sprintf("%v", r))
		}

	}

}
