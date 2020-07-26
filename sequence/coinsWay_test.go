package sequence

import (
	"fmt"
	"testing"
	"time"
)

func TestCoinsWay(t *testing.T) {
	arr1 := []int{10, 5, 1, 25}
	aim1 := 2000
	want := 1103021
	var r int
	t.Log(fmt.Sprintf("input coinsArr %v, aim: %d, want: %v", arr1, aim1, want))

	ts := time.Now()
	r = CoinsWay1(arr1, aim1)
	te := time.Now()
	fmt.Printf("CoinsWay1 cost %d us\n", te.Sub(ts).Nanoseconds())
	if r == want {
		t.Log("CoinsWay1 is ok")
	} else {
		t.Error("CoinsWay1 is not ok, result:", fmt.Sprintf("%v", r))
	}

	aim1 = 20000
	want = 1070270201
	t.Log(fmt.Sprintf("input coinsArr %v, aim: %d, want: %v", arr1, aim1, want))

	ts = time.Now()
	r = CoinsWay2(arr1, aim1)
	te = time.Now()

	fmt.Printf("CoinsWay2 cost %d us\n", te.Sub(ts).Nanoseconds())

	if r == want {
		t.Log("CoinsWay2 is ok")
	} else {
		t.Error("CoinsWay2 is not ok, result:", fmt.Sprintf("%v", r))
	}

	ts = time.Now()
	r = CoinsWay3(arr1, aim1)
	te = time.Now()

	fmt.Printf("CoinsWay3 cost %d us\n", te.Sub(ts).Nanoseconds())

	if r == want {
		t.Log("CoinsWay3 is ok")
	} else {
		t.Error("CoinsWay3 is not ok, result:", fmt.Sprintf("%v", r))
	}

	ts = time.Now()
	r = CoinsWay4(arr1, aim1)
	te = time.Now()

	fmt.Printf("CoinsWay4 cost %d us\n", te.Sub(ts).Nanoseconds())

	if r == want {
		t.Log("CoinsWay4 is ok")
	} else {
		t.Error("CoinsWay4 is not ok, result:", fmt.Sprintf("%v", r))
	}

	ts = time.Now()
	r = CoinsWay5(arr1, aim1)
	te = time.Now()

	fmt.Printf("CoinsWay5 cost %d us\n", te.Sub(ts).Nanoseconds())

	if r == want {
		t.Log("CoinsWay5 is ok")
	} else {
		t.Error("CoinsWay5 is not ok, result:", fmt.Sprintf("%v", r))
	}
}
