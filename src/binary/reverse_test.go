package binary

import (
	"testing"
	"fmt"
)

func TestReverse(t *testing.T) {
	var x uint32 = 9
	t.Log("input x:", x, fmt.Sprintf("(%b)", x))
	var want uint32 = 2415919104
	var r uint32 = ReverseBin32(x)
	if r == want {
		t.Log("reverse is ok:", want, fmt.Sprintf("(%b)", r))
	} else {
		t.Error("reverse is not ok:", r, fmt.Sprintf("(%b)", r), ",want:", want, fmt.Sprintf("(%b)", want))
	}
}
