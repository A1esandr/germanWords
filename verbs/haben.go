package main

import "germanWords/check"

func main() {

	w := [][2]string{
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

	check.Check(w)
}
