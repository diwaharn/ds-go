package util

import "fmt"

var DebugEnable bool = false

func Printf(format string, v ...interface{}) {
    if DebugEnable {
        fmt.Printf(format, v...)
    }
}

func Myabs(v int) int {
	if v < 0 {
		return v * -1
	}
	return v
}
