package sequence

import (
	"fmt"
)

func Hanoi(n int) {
	if n > 0 {
		do(n, "left", "mid", "right")
	}
}

func do(n int, from, by, to string) {
	if n == 1 {
		fmt.Printf("move from %s to %s by %s\n", from, to, by)
	} else {
		do(n-1, from, to, by)
		do(1, from, by, to)
		do(n-1, by, to, from)
	}
}
