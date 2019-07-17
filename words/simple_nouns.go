package main

import (
	"germanWords/check"
	"germanWords/reader"
)

func main() {
	w := reader.ReadCsv("csv/groups/simple_nouns.csv")
	check.Check(w)
}
