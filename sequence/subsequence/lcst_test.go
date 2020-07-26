package subsequence

import (
	"fmt"
	"testing"
)

func TestGetLcst(t *testing.T) {
	str1 := "ABC1234567DEFG"
	str2 := "HIJKL1234567MNOP"
	want := "1234567"
	t.Log(fmt.Sprintf("GetLcst, input str1 %v, str2 %v, want: %v", str1, str2, want))

	r := GetLcst(str1, str2)
	if fmt.Sprint(r) == fmt.Sprint(want) {
		t.Log("GetLcst is ok")
	} else {
		t.Error("GetLcst is not ok, result:", fmt.Sprintf("%v", r))
	}

	r = GetLcst2(str1, str2)
	if fmt.Sprint(r) == fmt.Sprint(want) {
		t.Log("GetLcst2 is ok")
	} else {
		t.Error("GetLcst2 is not ok, result:", fmt.Sprintf("%v", r))
	}
}
