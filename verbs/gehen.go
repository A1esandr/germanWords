package verbs

import (
	"germanWords/check"
	"germanWords/reader"
)

func Gehen() {
	w := reader.ReadCsv("csv/verbs/gehen.csv")
	check.Check(w)
}
