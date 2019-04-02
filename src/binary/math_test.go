package binary

import (
	"testing"
	"fmt"
)

func TestMin(t *testing.T) {
	var x, y = 4, 5
	var want = 4
	//min,max
	t.Log(fmt.Sprintf("input x=%d, y=%d, want min: %d", x, y, want))

	//power of 2
	fmt.Printf("%d is power of 2: %v\n", x, isPowerOf2(x))

	if min(x, y) == want {
		t.Log("min is ok")
	} else {
		t.Error("min is not ok")
	}
}

func TestMax(t *testing.T) {
	var x, y = 4, 5
	var want = 5
	//min,max
	t.Log(fmt.Sprintf("input x=%d, y=%d, want max: %d", x, y, want))

	//power of 2
	fmt.Printf("%d is power of 2: %v\n", x, isPowerOf2(x))

	if max(x, y) == want {
		t.Log("max is ok")
	} else {
		t.Error("max is not ok")
	}
}

func TestIsPowerOf2(t *testing.T) {
	var x = 4
	var want = true
	t.Log(fmt.Sprintf("input x=%d, is power of 2? want: %d", x, want))

	if isPowerOf2(x) == want {
		t.Log("isPowerOf2 is ok")
	} else {
		t.Error("isPowerOf2 is not ok")
	}
}
