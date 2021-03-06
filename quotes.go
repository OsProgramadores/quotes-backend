package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

type Quotes struct {
	Quote  string
	Author string
}

var quotacoes = map[string]Quotes{}

func main() {

	records, err := readData("quotes.csv")

	var i int = 0

	if err != nil {
		log.Fatal(err)
	}

	for _, record := range records {
		record := Quotes{
			Quote:  record[0],
			Author: record[1],
		}

		quotacoes[strconv.Itoa(i)] = record

		fmt.Printf("%v\n", quotacoes[strconv.Itoa(i)])
		fmt.Printf(("=>==========================================================================================\n"))

		i++
	}

	fmt.Printf("Processed %d quotes\n", i)

	router := mux.NewRouter()
	router.HandleFunc("/", quotesHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Assuming port# %s as standard", port)
	} else {
		log.Printf("Listening on port# %s", port)
	}

	log.Fatal(http.ListenAndServe(":"+port, router))
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

func quotesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var jsonResponse []byte

	var randomQuote = rand.Intn(16)

	fmt.Printf("Returning: %v\n", quotacoes[strconv.Itoa(randomQuote)])

	jsonResponse, err := json.Marshal(quotacoes[strconv.Itoa(randomQuote)])

	if err != nil {
		log.Printf("Error Marshalling: %v", err)
	}

	w.Write(jsonResponse)
}
