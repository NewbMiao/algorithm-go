package sequence

import (
	"testing"
	"fmt"
)

func TestDungeonGame(t *testing.T) {
	m := [][]int{{-2, -3, 3}, {-5, -10, 1}, {10, 30, -5},}
	want := 7
	t.Log(fmt.Sprintf("DungeonGame, input map info: %+v, want min blood init: %d", m, want))

	r := DungeonGame(m)
	if r == want {
		t.Log("DungeonGame is ok")
	} else {
		t.Error("DungeonGame is not ok, result:", fmt.Sprintf("%v", r))
	}

	r = DungeonGame2(m)
	if r == want {
		t.Log("DungeonGame2 is ok")
	} else {
		t.Error("DungeonGame2 is not ok, result:", fmt.Sprintf("%v", r))
	}
}
