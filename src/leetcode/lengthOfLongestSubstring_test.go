package leetcode

import (
	"fmt"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	d := []string{
		"abcabcbb",
		"a",
	}
	wants := []int{
		3,
		1,
	}
	for k, v := range d {
		t.Log(fmt.Sprintf("LengthOfLongestSubstring, input s %v, want: %v", v, wants[k]))
		r := lengthOfLongestSubstring(v)
		if r == wants[k] {
			t.Log("LengthOfLongestSubstring is ok")
		} else {
			t.Error("LengthOfLongestSubstring is not ok, result:", fmt.Sprintf("%v", r))
		}

	}

}
