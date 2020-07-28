package leetcode

import (
	"testing"
)

func TestConvert(t *testing.T) {
	type input struct {
		str string
		row int
	}
	inputs := []input{
		{"LEETCODEISHIRING", 3},
		{"LEETCODEISHIRING", 4},
		{"A", 2},
	}
	outputs := []string{
		"LCIRETOESIIGEDHN",
		"LDREOEIIECIHNTSG",
		"A",
	}
	for k, v := range inputs {
		expected := outputs[k]
		t.Logf("convert, input: %v", v)
		res := convert(v.str, v.row)
		if res == expected {
			t.Log("convert is ok")
		} else {
			t.Errorf("convert is not ok, result:%v, expect:%v", res, expected)
		}
		res = convertEven(v.str, v.row)
		if res == expected {
			t.Log("convert_even is ok")
		} else {
			t.Errorf("conconvert_evenvert is not ok, result:%v, expect:%v", res, expected)
		}
	}
}
