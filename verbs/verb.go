package verbs

import (
	"germanWords/check"
	"germanWords/reader"
)

func Verb(verb string) {
	w := reader.ReadCsv("csv/verbs/" + verb + ".csv")
	check.Check(w)
}
