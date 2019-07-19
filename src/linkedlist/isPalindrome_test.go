package linkedlist

import (
	"testing"
	"fmt"
	. "kit/linkedlist"

)

func TestIsPalindrome(t *testing.T) {
	l := NewList()
	want := false
	t.Log(fmt.Sprintf("isPalindrome, input list %v, want: %v", l, want))
	r := isPalindrome1(l)
	if r == want {
		t.Log("isPalindrome1 is ok")
	} else {
		t.Error("isPalindrome1 is not ok, result:", r)
	}
}
