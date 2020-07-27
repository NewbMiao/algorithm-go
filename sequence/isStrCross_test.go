package sequence

import (
	"fmt"
	"testing"
)

func TestIsStrCross(t *testing.T) {
	str1 := "1234"
	str2 := "abcd"
	aim := "1a23bcd4"
	want := true
	t.Log(fmt.Sprintf("IsStrCross, input str1 %v str2 %v aim %v, want: %v", str1, str2, aim, want))

	r := IsStrCross(str1, str2, aim)
	if r == want {
		t.Log("IsStrCross1 is ok")
	} else {
		t.Error("IsStrCross1 is not ok, result:", fmt.Sprintf("%v", r))
	}

	r = IsStrCross2(str1, str2, aim)
	if r == want {
		t.Log("IsStrCross2 is ok")
	} else {
		t.Error("IsStrCross2 is not ok, result:", fmt.Sprintf("%v", r))
	}
}
