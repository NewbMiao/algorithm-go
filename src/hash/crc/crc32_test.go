package crc

import (
	"testing"
	"fmt"
)

func TestCrc32(t *testing.T) {
	var data uint = 3
	var want uint = 192
	t.Log(fmt.Sprintf("input data=%d(%b), want crc32: %d(%b)", data, data, want, want))

	r := bitRev(data,8)
	if r == want {
		t.Log("crc32 is ok")
	} else {
		t.Error("crc32 is not ok, result:", fmt.Sprintf("%d(%b)", r, r))
	}
}
