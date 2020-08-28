package main

//#include <stdio.h>
import "C"

func testCgo() {
	C.puts(C.CString("Hello, World\n"))
}
