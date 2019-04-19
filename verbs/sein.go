package main

import "germanWords/check"

func main() {

	//Ic h(я) – bin (есть)
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
		{"I am","Ich bin"},
		{"You are (single)","Du bist"},
		{"He is","Er ist"},
		{"She is","Sie ist"},
		{"It is","Es ist"},
		{"We are","Wir sind"},
		{"You are (pl)","Ihr seid"},
		{"You are (polite)","Sie sind"},
		{"They are","Sie sind"},
	}

	check.Check(w)
}
