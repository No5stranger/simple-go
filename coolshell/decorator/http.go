package main

import (
	"fmt"
	"log"
	"net/http"
)

func WithServerHeader(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("---WithServerHeader")
		w.Header().Set("Server", "HelloServer")
		h(w, r)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	log.Println("Receive Request: %s from %s", r.URL.Path, r.RemoteAddr)
	fmt.Fprintf(w, "hello world!"+r.URL.Path)
}

func main() {
	http.HandleFunc("/v1/hello", WithServerHeader(hello))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Fatal("ListenAndServe: %s", err)
	}
}
