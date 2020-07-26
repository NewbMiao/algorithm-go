package mergesort

import (
	"strings"
	"testing"
)

func TestMergeSortRecursive(t *testing.T) {
	s := strings.Split("S O R T E X A M P L E", " ")
	t.Log("input array:", s)
	s = Sort(s)
	want := "A E E L M O P R S T X"
	if strings.Join(s, " ") == want {
		t.Log("merge sort is ok:", want)
	} else {
		t.Error("merge sort is not ok:", s, ",want:", want)
	}
}

func TestMergeSortRecursiveOnce(t *testing.T) {
	s := strings.Split("2 5 3 4 1 8 9 7 6 99", " ")
	t.Log("input array:", s)
	s = Sort(s)
	want := "1 2 3 4 5 6 7 8 9 99"
	if strings.Join(s, " ") == want {
		t.Log("merge sort is ok:", want)
	} else {
		t.Error("merge sort is not ok:", s, ",want:", want)
	}
}

func TestMergeSortBU(t *testing.T) {
	s := strings.Split("2 8 4 7 3 6 1 5 99 9", " ")
	t.Log("input array:", s)
	s = SortBU(s)
	want := "1 2 3 4 5 6 7 8 9 99"
	if strings.Join(s, " ") == want {
		t.Log("merge sort is ok:", want)
	} else {
		t.Error("merge sort is not ok:", s, ",want:", want)
	}
}
