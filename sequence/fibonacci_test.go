package sequence

import (
	"testing"
	"fmt"
)

func TestFibonacci(t *testing.T) {
	n := 20
	want := 6765
	t.Log(fmt.Sprintf("Fibonacci3's %d is, want: %v", n, want))

	fmt.Println(Fibonacci1(n), Fibonacci2(n))
	r := Fibonacci3(n)
	if r == want {
		t.Log("Fibonacci3 is ok")
	} else {
		t.Error("Fibonacci3 is not ok, result:", fmt.Sprintf("%v", r))
	}
}
