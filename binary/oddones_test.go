package binary

import (
	"fmt"
	"testing"
)

func TestOddOnes(t *testing.T) {
	var s uint = 98734520
	want := true
	t.Log(fmt.Sprintf("%d(0%b) has odd one? %v", s, s, want))
	oddBools := []bool{
		judgeBinOddOne1(s), //自身移位压缩奇偶信息
		judgeBinOddOne2(s), //x-1依次消除1
		judgeBinOddOne3(s), //查表法：256区间内1的个数表
		judgeBinOddOne4(s), //区间压缩奇偶信息
		judgeBinOddOne5(s), //查表法：256区间内1的奇偶表
		judgeBinOddOne6(s), //并行查奇偶表
		judgeBinOddOne7(s), //乘法 8 operations
	}
	for k, v := range oddBools {
		if v == want {
			t.Log(fmt.Sprintf("judgeBinOddOne%d is ok", k+1))
		} else {
			t.Errorf("judgeBinOddOne%d is not ok", k+1)
		}
	}
}
