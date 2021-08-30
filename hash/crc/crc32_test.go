package crc

import (
	"fmt"
	"hash/crc32"
	"testing"
)

func TestCrc32(t *testing.T) {
	data := "1234567890"
	var want uint32 = 639479525
	t.Log(fmt.Sprintf("input data=%s, want crc32: %d(%b)", data, want, want))

	fmt.Println("hash/crc32 result is:", crc32.ChecksumIEEE([]byte(data)))
	r := Crc32([]byte(data))
	if r == want {
		t.Log("crc32 is ok")
	} else {
		t.Error("crc32 is not ok, result:", fmt.Sprintf("%d(%b)", r, r))
	}
}
