package app

import (
	"bufio"
	"flag"
	"fmt"
	"germanWords/verbs"
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
	if mode == "web" {
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
	fmt.Println("Select exercise: numbers (0), pronoun (1), nouns (2), verbs (3)")
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	switch text {
	case "0":
		fmt.Println("Select exercise: 0-9 (0), 10-100 (1)")
		digit, _ := reader.ReadString('\n')
		digit = strings.Replace(digit, "\n", "", -1)
		switch digit {
		case "0":
			words.Digits()
		case "1":
			words.Numbers()
		default:
			fmt.Println("Wrong exercise select. Exit.")
		}
	case "1":
		words.Pronoun()
	case "2":
		words.Nouns()
	case "3":
		fmt.Println("Select verb: gehen (0), haben (1), kommen (2), sein (3), wohnen (4)")
		verb, _ := reader.ReadString('\n')
		verb = strings.Replace(verb, "\n", "", -1)
		switch verb {
		case "0":
			verbs.Gehen()
		case "1":
			verbs.Haben()
		case "2":
			verbs.Kommen()
		case "3":
			verbs.Sein()
		case "4":
			verbs.Wohnen()
		default:
			fmt.Println("Wrong verb select. Exit.")
		}
	default:
		fmt.Println("Wrong select. Exit.")
	}
}
