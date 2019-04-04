package crc

import (
	"testing"
	"fmt"
)

func TestCrc4(t *testing.T) {
	var data uint = 0x035B
	var want uint = 0x0000000E
	t.Log(fmt.Sprintf("input data=%d(%b), want crc4: %d(%b)", data, data, want, want))

	r := Crc4(data)
	if r == want {
		t.Log("crc4 is ok")
	} else {
		t.Error("crc4 is not ok, result:", fmt.Sprintf("%d(%b)", r, r))
	}
}
