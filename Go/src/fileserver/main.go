package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"runtime"
	"runtime/pprof"
)

var MyMux = &myHTTPServerMux{muxHandlers: http.NewServeMux()}

type myHTTPServerMux struct {
	muxHandlers *http.ServeMux
}

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

func (mux *myHTTPServerMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	var path = r.URL.Path
	handler, pattern := mux.muxHandlers.Handler(r)
	log.Printf("path:%s\n", path)
	log.Println("http pattern:", pattern)
	if handler == nil || pattern == "" {
		log.Printf("no handler for Request, path:%s\n", path)
		w.Write([]byte("hell world"))
		return
	}
	handler.ServeHTTP(w, r)
}

func main() {
	MyMux.muxHandlers.Handle("/", http.FileServer(http.Dir("D:\\temp\\download")))
	s := &http.Server{
		Addr:    "172.18.4.237:8090",
		Handler: MyMux,
		// ReadTimeout:    10 * time.Second,
		//WriteTimeout:   120 * time.Second,
		MaxHeaderBytes: 1 << 8,
	}
	s.ListenAndServe()
	waitInput()
}

func waitInput() {
	var cmd string
	for {
		_, err := fmt.Scanf("%s\n", &cmd)
		if err != nil {
			//log.Println("Scanf err:", err)
			continue
		}

		switch cmd {
		case "exit", "quit":
			log.Println("exit by user")
			return
		case "gr":
			log.Println("current goroutine count:", runtime.NumGoroutine())
			break
		case "gd":
			pprof.Lookup("goroutine").WriteTo(os.Stdout, 1)
			break
		default:
			break
		}
	}
}
