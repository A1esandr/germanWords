package app

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

	s := Read(name)

	r := csv.NewReader(strings.NewReader(s))
	r.Comma = rune([]byte(";")[0])

	var res [][]string

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		res = append(res, record)
	}

	return res
}

func Read(name string) string {
	file, err := os.Open(name)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	// defer the closing of our csvFile so that we can parse it later on
	defer func() {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	byteArray, _ := ioutil.ReadAll(file)

	return string(byteArray[:])
}
