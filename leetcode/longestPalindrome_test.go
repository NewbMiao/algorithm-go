package leetcode

import (
	"fmt"
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	d := []string{
		"babad",
		"cbbd",
		"aaaa",
	}
	wants := []string{
		"aba",
		"bb",
		"aaaa",
	}
	wants2 := []string{
		"bab",
		"bb",
		"aaaa",
	}
	for k, v := range d {
		str := fmt.Sprintf("LongestPalindrome, input s %v, want: %v", v, wants[k])
		if wants[k] != wants2[k] {
			str = fmt.Sprintf("%s or %s", str, wants2[k])
		}
		t.Log(str)
		r := longestPalindrome(v)
		if r == wants[k] || r == wants2[k] {
			t.Log("LongestPalindrome is ok")
		} else {
			t.Error("LongestPalindrome is not ok, result:", fmt.Sprintf("%v", r))
		}

		r = longestPalindrome1(v)
		if r == wants[k] || r == wants2[k] {
			t.Log("LongestPalindrome1 is ok")
		} else {
			t.Error("LongestPalindrome1 is not ok, result:", fmt.Sprintf("%v", r))
		}

		r = longestPalindrome2(v)
		if r == wants[k] || r == wants2[k] {
			t.Log("LongestPalindrome2 is ok")
		} else {
			t.Error("LongestPalindrome2 is not ok, result:", fmt.Sprintf("%v", r))
		}
	}
}
