package stringalg

import (
	"fmt"
	"testing"
)

func TestStrstr(t *testing.T) {
	str := "abcabcababaccc"
	match := "ababa"
	want := 6
	t.Log(fmt.Sprintf("sunday: input %s, strStr index of %s, want: %v", str, match, want))

	r := strStr(str, match)
	if r == want {
		t.Log("strStr is ok")
	} else {
		t.Error("strStr is not ok, result:", fmt.Sprintf("%v", r))
	}
}
