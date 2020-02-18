package string_alg

import (
	"testing"
	"fmt"
)

func TestGetIndexOf(t *testing.T) {
	str := "abcabcababaccc"
	match := "ababa"
	want := 6
	t.Log(fmt.Sprintf("KMP  input %s, GetIndexOf %s, want: %v", str, match, want))

	r := GetIndexOf(str, match)
	if r == want {
		t.Log("GetIndexOf is ok")
	} else {
		t.Error("GetIndexOf is not ok, result:", fmt.Sprintf("%v", r))
	}
}
