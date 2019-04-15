package main

import (
	"germanWords/check"
)

func main() {

	//ich -я
	//
	//du - ты
	//
	//er, sie, es - он, она, оно
	//
	//Plural(мн.ч.)
	//
	//wir -мы
	//
	//ihr - вы
	//
	//sie, Sie* — они, Вы

	w := [][2]string{
		{"I","ich"},
		{"You (single)","du"},
		{"You (sev)","ihr"},
		{"You (sev, polite)","Sie"},
		{"He","er"},
		{"She","sie"},
		{"It","es"},
		{"We","wir"},
		{"They","sie"},
	}

	check.Check(w)
}
