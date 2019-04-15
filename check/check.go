package check

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Check(s [][2]string) {
	reader := bufio.NewReader(os.Stdin)

	for i:=0; i< len(s); i++ {
		subCheck(s[i], *reader)
	}
}

func subCheck(p [2]string, reader bufio.Reader) {
	fmt.Print(p[0]+": ")
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)

	if text == p[1] {
		fmt.Println("success")
	}
}

func PrintAll(w [9][2]string) {
	s := ""
	for i:=0; i< len(w); i++ {
		s += w[i][1] + ", "
	}
	s += "\n"
	fmt.Print(s)
}

func printBytes(s string) {
	for i:= 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
}