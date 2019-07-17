package main

import (
	"encoding/csv"
	"fmt"
	"germanWords/reader"
	"io"
	"log"
	"strings"
)

func main() {
	s:=reader.Read("csv/groups/simple_nouns.csv")
	fmt.Print(s)

	r := csv.NewReader(strings.NewReader(s))

	var res [][]string

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		ss:=strings.Split(record[0], ";")
		ss=ss[:2]
		res = append(res, ss)
	}
}
