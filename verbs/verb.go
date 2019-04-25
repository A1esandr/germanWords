package main

import (
	"germanWords/check"
	"germanWords/reader"
	"os"
)

func main() {
	argsWithoutProg := os.Args[1:]
	for i := 0; i < len(argsWithoutProg); i++ {
		w := reader.ReadCsv("csv/verbs/" + argsWithoutProg[i] + ".csv")
		check.Check(w)
	}
}
