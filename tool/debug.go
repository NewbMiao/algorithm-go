package tool

import (
	"fmt"
	"flag"
	"testing"
)
// https://github.com/golang/go/issues/31859#issuecomment-489889428
var _ = func() bool {
	testing.Init()
	return true
}()

var DEBUG bool

func init() {
	flag.BoolVar(&DEBUG, "debug", false, "show trace log")
	flag.Parse()
}
func Trace(fmtStr string, args ... interface{}) {
	if DEBUG {
		fmt.Printf("[Trace] "+fmtStr, args...)
	}
}
