package main

import (
	"os"
	"runtime/trace"
)

func testTrace() {
	trace.Start(os.Stderr)
	defer trace.Stop()

	ch := make(chan string)
	go func() {
		ch <- "EDDYCJY"
	}()

	<-ch
}
