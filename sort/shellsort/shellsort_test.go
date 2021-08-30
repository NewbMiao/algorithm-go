package shellsort

import (
	"strings"
	"testing"
)

type Sitem []string

func (s Sitem) Len() int {
	return len(s)
}

func (s Sitem) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s Sitem) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func TestShellSort(t *testing.T) {
	s := strings.Split("S O R T E X A M P L E", " ")
	t.Log("input array:", s)
	Sort(Sitem(s))
	want := "A E E L M O P R S T X"
	if strings.Join(s, " ") == want {
		t.Log("shell sort is ok:", want)
	} else {
		t.Error("shell sort is not ok:", s, ",want:", want)
	}
}
