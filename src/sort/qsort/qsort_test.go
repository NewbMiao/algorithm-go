package qsort

import (
	"testing"
	"strings"
)

func TestQSort(t *testing.T) {

	s := strings.Split("S O R T E X A M P L E", " ")
	t.Log("input array:", s)
	Sort(Sitem(s))
	want := "A E E L M O P R S T X"
	if strings.Join(s, " ") == want {
		t.Log("qsort sort is ok:", want)
	} else {
		t.Error("qsort sort is not ok:", s, ",want:", want)
	}
}

func TestQSortX(t *testing.T) {

	s := strings.Split("S O R T E X A M P L E", " ")
	t.Log("input array:", s)
	SortX(s)
	want := "A E E L M O P R S T X"
	if strings.Join(s, " ") == want {
		t.Log("qsort sort is ok:", want)
	} else {
		t.Error("qsort sort is not ok:", s, ",want:", want)
	}
}

func TestQSort3Way(t *testing.T) {

	s := strings.Split("S S O R T S O R T E X A M P L E E", " ")
	t.Log("input array:", s)
	Sort3way(s)
	want := "A E E E L M O O P R R S S S T T X"
	if strings.Join(s, " ") == want {
		t.Log("qsort sort is ok:", want)
	} else {
		t.Error("qsort sort is not ok:", s, ",want:", want)
	}
}

func TestQSort3Way2(t *testing.T) {

	s := strings.Split("2 1 1 1 2 3 1 2 3 1 2 2 1 1 1 2 3 1 2 3 1 2 3 3 6 6 4 4 5 5 5 3 3 2 1 1 1 2 3 1 2 3 1 2 3 3 6 6 4 4 5 5 5 6 6 4 4 5 5 5", " ")
	t.Log("input array:", s)
	Sort3way(s)
	want := "1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 2 2 2 2 2 2 2 2 2 2 2 2 3 3 3 3 3 3 3 3 3 3 3 3 4 4 4 4 4 4 5 5 5 5 5 5 5 5 5 6 6 6 6 6 6"
	if strings.Join(s, " ") == want {
		t.Log("qsort sort is ok:", want)
	} else {
		t.Error("qsort sort is not ok:", s, ",want:", want)
	}
}
