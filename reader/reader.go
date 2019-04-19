package reader

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func ReadCsv(name string) [][]string {
	csvFile, err := os.Open(name)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	// defer the closing of our csvFile so that we can parse it later on
	defer csvFile.Close()

	byteArray, _ := ioutil.ReadAll(csvFile)

	s := string(byteArray[:])

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

	return res
}
