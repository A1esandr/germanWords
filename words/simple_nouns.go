package words

import (
	"germanWords/check"
	"germanWords/reader"
)

func Nouns() {
	w := reader.ReadCsv("csv/groups/simple_nouns.csv")
	check.Check(w)
}
