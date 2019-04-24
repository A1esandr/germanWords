package main

import (
	"germanWords/check"
	"germanWords/pars"
)

func main() {
	s := `//0 – null	10 – zehn	20 – zwanzig	10 – zehn
	//1 – eins	11 – elf	21 – einundzwanzig (1 и 20)	20 – zwanzig
	//2 – zwei	12 – zwölf	22 – zweiundzwanzig (2 и 20)	30 – dreißig
	//3 – drei	13 – dreizehn (3,10)	23 – dreiundzwanzig (3 и 20)	40 – vierzig
	//4 – vier	14 – vierzehn (4,10)	24 – vierundzwanzig	50 – fünfzig
	//5 – fünf	15 – fünfzehn (5,10)	25 – fünfundzwanzig	60 – sechzig
	//6 – sechs	16 – sechzehn	26 – sechsundzwanzig	70 – siebzig
	//7 – sieben	17 – siebzehn	27 – siebenundzwanzig	80 – achtzig
	//8 – acht	18 – achtzehn	28 – achtundzwanzig	90 – neunzig
	//9 – neun	19 – neunzehn	29 – neunundzwanzig	100 – hundert`

	ss := pars.Parse(s)

	check.Check(ss)
}
