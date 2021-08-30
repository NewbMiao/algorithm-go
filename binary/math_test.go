package binary

import (
	"fmt"
	"testing"
)

func TestMin(t *testing.T) {
	x, y := 4, 5
	want := 4
	t.Log(fmt.Sprintf("input x=%d, y=%d, want min: %d", x, y, want))

	if Min(x, y) == want {
		t.Log("min is ok")
	} else {
		t.Error("min is not ok")
	}
}

func TestMax(t *testing.T) {
	x, y := 4, 5
	want := 5
	t.Log(fmt.Sprintf("input x=%d, y=%d, want max: %d", x, y, want))
	if Max(x, y) == want {
		t.Log("max is ok")
	} else {
		t.Error("max is not ok")
	}
}

func TestIsPowerOf2(t *testing.T) {
	x := 4
	want := true
	t.Log(fmt.Sprintf("input x=%d, is power of 2? want: %v", x, want))

	if IsPowerOf2(x) == want {
		t.Log("isPowerOf2 is ok")
	} else {
		t.Error("isPowerOf2 is not ok")
	}
}

func TestMPowerOfN(t *testing.T) {
	m, n := 5, 6
	want := 15625
	t.Log(fmt.Sprintf("input: %d power of %d? want: %d", m, n, want))

	if mPowerOfN(m, n) == want {
		t.Log("mPowerOfN is ok")
	} else {
		t.Error("mPowerOfN is not ok, get:", mPowerOfN(m, n))
	}
}
