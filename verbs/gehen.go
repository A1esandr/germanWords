package main

import (
	"germanWords/check"
	"germanWords/reader"
)

func main() {
	w:=reader.ReadCsv("csv/verbs/gehen.csv")
	check.Check(w)
}
