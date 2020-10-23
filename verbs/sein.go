package verbs

import "germanWords/check"

func Sein() {

	//Ich(я) – bin (есть)
	//
	//Du (ты) – bist (есть)
	//
	//Er/sie/es(он/она/оно) — ist (есть)
	//
	//Plural (множественное число)
	//
	//Wir(мы) — sind (есть)
	//
	//Ihr(вы) — seid (есть)
	//
	//Sie/sie(Вы/они) — sind (есть)

	w := [][]string{
		{"I am", "ich bin"},
		{"You are (single)", "du bist"},
		{"He is", "er ist"},
		{"She is", "sie ist"},
		{"It is", "es ist"},
		{"We are", "wr sind"},
		{"You are (pl)", "ihr seid"},
		{"You are (polite)", "Sie sind"},
		{"They are", "sie sind"},
	}

	check.Check(w)
}
