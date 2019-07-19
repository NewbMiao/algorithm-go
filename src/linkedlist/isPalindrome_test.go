package linkedlist

import (
	. "kit/linkedlist"
	"testing"
	"fmt"
)

func TestIsPalindrome(t *testing.T) {

	ds := [][]interface{}{
		{},
		{1},
		{1, 2},
		{1, 1},
		{1, 2, 3},
		{1, 2, 1},
		{1, 2, 2, 1},
		{1, 2, 3, 2, 1},
	}
	ws := []bool{
		true,
		true,
		false,
		true,
		false,
		true,
		true,
		true,
	}
	for k := range ds {
		t.Log(fmt.Sprintf("isPalindrome, input list %v, want: %v", ds[k], ws[k]))
		l := Arr2LList(ds[k])
		var r bool
		r = isPalindrome1(l)
		if r == ws[k] {
			t.Log("isPalindrome1 is ok")
		} else {
			t.Error("isPalindrome1 is not ok, result:", r)
		}

		r = isPalindrome2_1(l)
		if r == ws[k] {
			t.Log("isPalindrome2_1 is ok")
		} else {
			t.Error("isPalindrome2_1 is not ok, result:", r)
		}
		r = isPalindrome2_2(l)
		if r == ws[k] {
			t.Log("isPalindrome2_2 is ok")
		} else {
			t.Error("isPalindrome2_2 is not ok, result:", r)
		}

		r = isPalindrome3(l)
		if r == ws[k] {
			t.Log("isPalindrome3 is ok")
		} else {
			t.Error("isPalindrome3 is not ok, result:", r)
		}
	}

}
