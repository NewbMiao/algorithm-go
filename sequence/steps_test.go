package sequence

import (
	"fmt"
	"testing"
)

func TestSteps(t *testing.T) {
	n := 20
	want := 10946
	t.Log(fmt.Sprintf("%d Steps, want: %v", n, want))

	fmt.Println(Steps1(n), Steps2(n))
	r := Steps3(n)
	if r == want {
		t.Log("Steps is ok")
	} else {
		t.Error("Steps is not ok, result:", fmt.Sprintf("%v", r))
	}
}
