package binary

import (
	"fmt"
	"testing"
)

func TestSwap(t *testing.T) {
	// swap bit seq
	var i, j, n uint = 1, 5, 3
	var bInt uint = 47
	var want uint = 227
	t.Log(fmt.Sprintf("%d(%b) swap left %d bit from %d and %d, want: %d(%b)\n",
		bInt, bInt, n, i, j, want, want))
	rInt := swapBitSeq(i, j, n, bInt)

	if rInt == want {
		t.Log("swapBitSeq is ok")
	} else {
		t.Error("swapBitSeq is not ok", rInt, fmt.Sprintf("(%b)", rInt))
	}
}
