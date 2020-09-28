package app

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
)

type App struct {
}

func (a *App) Start() {
	flag.Parse()
	fmt.Println("v0.1")
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}

	http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
