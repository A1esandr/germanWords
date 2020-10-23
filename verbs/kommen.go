package verbs

import (
	"germanWords/check"
	"germanWords/reader"
)

func Kommen() {
	w := reader.ReadCsv("csv/verbs/kommen.csv")
	check.Check(w)
}
