package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	flag.Parse()
	fmt.Println("v0.1")
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}

	http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
