package verbs

import (
	"germanWords/check"
	"germanWords/reader"
)

func Wohnen() {
	w := reader.ReadCsv("csv/verbs/wohnen.csv")
	check.Check(w)
}
