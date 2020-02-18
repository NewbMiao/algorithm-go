package sequence

import (
	"testing"
	"fmt"
)

func TestMinCoinsOne(t *testing.T) {
	arr1 := []int{100, 20, 5, 10, 2, 50, 1};
	aim1 := 17019;
	want := 174
	var r int
	t.Log(fmt.Sprintf("input coinsArr %v, aim: %d, want: %v", arr1, aim1, want))

	r = MinCoins1(arr1, aim1)
	if r == want {
		t.Log("MinCoins1 is ok")
	} else {
		t.Error("MinCoins1 is not ok, result:", fmt.Sprintf("%v", r))
	}

	r = MinCoins2(arr1, aim1)
	if r == want {
		t.Log("MinCoins2 is ok")
	} else {
		t.Error("MinCoins2 is not ok, result:", fmt.Sprintf("%v", r))
	}
}

func TestMinCoinsTwo(t *testing.T) {
	arr2 := []int{10, 100, 2, 5, 5, 5, 10, 1, 1, 1, 2, 100};
	aim2 := 223;
	want := 6
	var r int
	t.Log(fmt.Sprintf("input coinsArr2 %v, aim: %d, want: %v", arr2, aim2, want))

	r = MinCoins3(arr2, aim2)
	if r == want {
		t.Log("MinCoins3 is ok")
	} else {
		t.Error("MinCoins3 is not ok, result:", fmt.Sprintf("%v", r))
	}

	r = MinCoins4(arr2, aim2)
	if r == want {
		t.Log("MinCoins4 is ok")
	} else {
		t.Error("MinCoins4 is not ok, result:", fmt.Sprintf("%v", r))
	}
}
