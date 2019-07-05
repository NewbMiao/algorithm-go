package leetcode

import (
	"testing"
	"fmt"
)

func TestIsValid(t *testing.T) {
	ds := []string{
		"()",
		"()[]{}",
		"(]",
		"([)]",
		"{[]}",
	}

	ws := []bool{
		true,
		true,
		false,
		false,
		true,
	}
	for k, v := range ds {
		t.Log(fmt.Sprintf("isValid, input str %v, want: %v", v, ws[k]))
		r := isValid(v)
		if r == ws[k] {
			t.Log("isValid is ok")
		} else {
			t.Error("isValid is not ok, result:", r)
		}

	}
}
