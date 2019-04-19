package main

import (
	"germanWords/check"
	"germanWords/reader"
)

func main() {
	w:=reader.ReadCsv("csv/verbs/wohnen.csv")
	check.Check(w)
}
