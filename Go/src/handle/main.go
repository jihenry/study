package main

import (
	"fmt"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello dddworld")
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("D:\\temp")))
	http.ListenAndServe("127.0.0.1:8000", nil)
	// http.HandleFunc("/", IndexHandler)
	// http.ListenAndServe(":8000", nil)
}
