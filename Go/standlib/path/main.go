package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(getAppPath(os.Args[0]))
}
