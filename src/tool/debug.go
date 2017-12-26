package tool

import (
	"fmt"
	"flag"
)

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
