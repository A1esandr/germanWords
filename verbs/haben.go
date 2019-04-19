package main

import (
	"germanWords/check"
	"germanWords/shuffle"
)

func main() {

	w := [][]string{
		{"I have","ich habe"},
		{"He has","er hat"},
		{"She has","sie hat"},
		{"It has","es hat"},
		{"We have","wir haben"},
		{"They have","sie haben"},
		{"You have (polite)","Sie haben"},
		{"You have (single)","du hast"},
		{"You have (pl)","ihr habt"},
	}

	shuffle.Shuffle(w)

	check.Check(w)
}
