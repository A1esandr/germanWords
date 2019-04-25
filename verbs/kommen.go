package main

import (
	"germanWords/check"
	"germanWords/reader"
)

func main() {
	w := reader.ReadCsv("csv/verbs/kommen.csv")
	check.Check(w)
}
