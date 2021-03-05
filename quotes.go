package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type Quotes struct {
	quote  string
	author string
}

func main() {

	records, err := readData("quotes.csv")

	if err != nil {
		log.Fatal(err)
	}

	for _, record := range records {
		record := Quotes{
			quote:  record[0],
			author: record[1],
		}

		fmt.Printf("%s by %s\n", record.quote, record.author)
		fmt.Print(("=>==========================================================================================\n"))
	}
}

func readData(fileName string) ([][]string, error) {

	f, err := os.Open(fileName)

	if err != nil {
		return [][]string{}, err
	}
	defer f.Close()

	r := csv.NewReader(f)

	// skip first line
	if _, err := r.Read(); err != nil {
		return [][]string{}, err
	}

	records, err := r.ReadAll()

	if err != nil {
		return [][]string{}, err
	}

	return records, nil
}
