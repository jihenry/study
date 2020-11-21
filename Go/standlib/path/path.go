package main

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func getAppPath(file string) string {
	fmt.Println("file:", file)
	path, _ := exec.LookPath(file)
	fmt.Println("path:", path)
	abs, _ := filepath.Abs(path)
	fmt.Println("abs:", abs)
	var a net.TCPConn
	index := strings.LastIndex(abs, string(os.PathSeparator))
	return abs[:index]
}
