package main

import (
	"flag"
	"fmt"
)

func main() {
	s := "xxxxxx"
	flag.Parse()
	// pprof.StartCPUProfile()
	fmt.Println("s:", s)
	// pprof.StopCPUProfile()
}
