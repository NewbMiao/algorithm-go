package leetcode

import (
	"testing"
)

func TestMaxProfit(t *testing.T) {
	inputs := [][]int{
		{7, 1, 5, 3, 6, 4},
		{1, 2, 3, 4, 5},
		{7, 6, 4, 3, 1},
	}
	outputs := []int{
		7,
		4,
		0,
	}
	for k, v := range inputs {
		expected := outputs[k]
		t.Logf("maxProfit, input: %v, expect: %v", v, expected)
		res := maxProfit(v)
		if res == expected {
			t.Log("maxProfit is ok")
		} else {
			t.Errorf("maxProfit is not ok, result:%v, expect:%v", res, expected)
		}

		res = maxprofitDp(v)
		if res == expected {
			t.Log("maxProfit_dp is ok")
		} else {
			t.Errorf("maxProfit_dp is not ok, result:%v, expect:%v", res, expected)
		}
	}
}
