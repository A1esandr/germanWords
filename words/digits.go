package words

import (
	"fmt"
	"germanWords/check"
	"germanWords/reader"
)

func Digits() {
	fmt.Println("Digits exercise. Type given digit in german.")
	w := reader.ReadCsv("csv/groups/digits.csv")
	check.Check(w)
}

func Numbers() {
	fmt.Println("Numbers exercise. Type given number in german.")
	w := reader.ReadCsv("csv/groups/numbers.csv")
	check.Check(w)
}
