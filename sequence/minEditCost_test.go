package sequence

import (
	"fmt"
	"testing"
)

func TestMinEditCost(t *testing.T) {
	type tData struct {
		str1, str2       string
		ic, dc, rc, want int
	}
	wants := []tData{
		{
			"ab12cd3",
			"abcdf",
			5, 3, 2, 8},
		{
			"abcdf",
			"ab12cd3",
			3, 2, 4, 10},
		{
			"",
			"ab12cd3",
			1, 7, 5, 7},
		{
			"abcdf",
			"",
			2, 9, 8, 45},
	}
	for _, v := range wants {
		var r int
		t.Log(fmt.Sprintf("minEditCost, input edit info: %+v", v))

		r = MinEditCost(v.str1, v.str2, v.dc, v.rc, v.ic)
		if r == v.want {
			t.Log("minEditCost1 is ok")
		} else {
			t.Error("minEditCost1 is not ok, result:", fmt.Sprintf("%v", r))
		}

		r = MinEditCost2(v.str1, v.str2, v.dc, v.rc, v.ic)
		if r == v.want {
			t.Log("minEditCost2 is ok")
		} else {
			t.Error("minEditCost2 is not ok, result:", fmt.Sprintf("%v", r))
		}
	}
}
