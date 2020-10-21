package app

import (
	"bufio"
	"flag"
	"fmt"
	"germanWords/words"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type App struct {
}

func (a *App) Start() {
	flag.Parse()
	fmt.Println("German words program")
	fmt.Println("v0.1")

	mode := os.Getenv("MODE")
	if mode != "cli" {
		helloHandler := func(w http.ResponseWriter, req *http.Request) {
			_, err := io.WriteString(w, "Hello, world!\n")
			if err != nil {
				fmt.Println(err.Error())
			}
		}

		http.HandleFunc("/hello", helloHandler)
		log.Fatal(http.ListenAndServe(":8081", nil))
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Select exercise: digits (0), pronoun (1), nouns (2)")
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	switch text {
	case "0":
		words.Digits()
	case "1":
		words.Pronoun()
	case "2":
		words.Nouns()
	default:
		fmt.Println("Wrong select. Exit.")
	}
}
