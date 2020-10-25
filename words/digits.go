package words

import (
	"fmt"
	"germanWords/check"
	"germanWords/pars"
)

func Digits() {
	s := `//0 – null	
	//1 – eins
	//2 – zwei
	//3 – drei
	//4 – vier
	//5 – fünf
	//6 – sechs
	//7 – sieben	
	//8 – acht
	//9 – neun`

	ss := pars.Parse(s)

	fmt.Println("Digits exercise. Type given digit in german.")
	check.Check(ss)
}

func Tens() {
	s := `//10 – zehn	20 – zwanzig	10 – zehn
	//11 – elf	21 – einundzwanzig (1 и 20)	20 – zwanzig
	//12 – zwölf	22 – zweiundzwanzig (2 и 20)	30 – dreißig
	//13 – dreizehn (3,10)	23 – dreiundzwanzig (3 и 20)	40 – vierzig
	//14 – vierzehn (4,10)	24 – vierundzwanzig	50 – fünfzig
	//15 – fünfzehn (5,10)	25 – fünfundzwanzig	60 – sechzig
	//16 – sechzehn	26 – sechsundzwanzig	70 – siebzig
	//17 – siebzehn	27 – siebenundzwanzig	80 – achtzig
	//18 – achtzehn	28 – achtundzwanzig	90 – neunzig
	//19 – neunzehn	29 – neunundzwanzig	100 – hundert`

	ss := pars.Parse(s)

	fmt.Println("Digits exercise. Type given digit in german.")
	check.Check(ss)
}
