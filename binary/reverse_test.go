package binary

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	var x uint32 = 13
	var want uint32 = 2952790016

	t.Log("input 32bit x:", x, fmt.Sprintf("(%b) \nreverse is %d(%b)", x, want, want))
	var r uint32 = ReverseBin32(x)
	if r == want {
		t.Log("reverseBin32 is ok")
	} else {
		t.Error("reverseBin32 is not ok", r, fmt.Sprintf("(%b)", r))
	}

	var oInt uint = 47
	var want2 uint = 17582052945254416384
	t.Log("input 64bit x:", x, fmt.Sprintf("(%b) \nreverse is %d(%b)", oInt, want2, want2))
	var reverseInts = []uint{
		reverseBit1(oInt),
		reverseBit2(oInt),
	}
	for k, v := range reverseInts {
		if v == want2 {
			t.Log(fmt.Sprintf("reverseBit%d is ok", k+1))
		} else {
			t.Error(fmt.Sprintf("reverseBit%d is not ok", k+1), v, fmt.Sprintf("(%b)", v))
		}
	}
}
