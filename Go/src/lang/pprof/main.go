package main

import (
	"lang/pprof/data"
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	go func() {
		for {
			log.Println(data.Add("https://github.com/EDDYCJY"))
		}
	}()
	http.ListenAndServe("0.0.0.0:6060", nil)
}
