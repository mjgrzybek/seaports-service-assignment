package safeExit

var Done chan struct{}

func init() {
	Done = make(chan struct{})
}
