package subsequence

import (
	"fmt"
	"testing"
)

func TestGetLcse(t *testing.T) {
	str1 := "A1BC2D3EFGH45I6JK7LMN"
	str2 := "12OPQ3RST4U5V6W7XYZ"

	want := "1234567"
	t.Log(fmt.Sprintf("GetLcse, input str1 %v, str2 %v, want: %v", str1, str2, want))

	r := GetLcse(str1, str2)
	if fmt.Sprint(r) == fmt.Sprint(want) {
		t.Log("GetLcse is ok")
	} else {
		t.Error("GetLcse is not ok, result:", fmt.Sprintf("%v", r))
	}
}
