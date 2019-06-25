package sequence

import (
	"testing"
	"fmt"
)

func TestCows(t *testing.T) {
	n := 20
	want := 1873
	t.Log(fmt.Sprintf("after %d years, want: %v cows", n, want))

	fmt.Println("cow1:", Cow1(n), "cow2:", Cow2(n))
	r := Cow3(n)
	if r == want {
		t.Log("Cows is ok")
	} else {
		t.Error("Cows is not ok, result:", fmt.Sprintf("%v", r))
	}
}
