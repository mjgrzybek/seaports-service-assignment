package safeExit

import "fmt"

var Done chan struct{}

func init() {
	Done = make(chan struct{})
}

func MarkDone() {
	Done <- struct{}{}
}
