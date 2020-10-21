package words

import (
	"germanWords/check"
)

func Pronoun() {

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

	w := [][]string{
		{"I", "ich"},
		{"You (single)", "du"},
		{"You (sev)", "ihr"},
		{"You (sev, polite)", "Sie"},
		{"He", "er"},
		{"She", "sie"},
		{"It", "es"},
		{"We", "wir"},
		{"They", "sie"},
	}

	check.Check(w)
}
