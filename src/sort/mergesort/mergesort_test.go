package mergesort

import (
	"testing"
	"strings"
)

func TestMergeSortRecursive(t *testing.T) {
	s := strings.Split("S O R T E X A M P L E", " ")
	t.Log("input array:", s)
	s = Sort(s)
	want := "A E E L M O P R S T X"
	if strings.Join(s, " ") == want {
		t.Log("shell sort is ok:", want)
	} else {
		t.Error("shell sort is not ok:", s, ",want:", want)
	}
}
