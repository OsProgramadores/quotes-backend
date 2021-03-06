package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type singleQuote struct {
	Quote  string
	Author string
}

type quotes struct {
	quotacoes []singleQuote
}

func (x *quotes) Handler(w http.ResponseWriter, r *http.Request) {
	var jsonResponse []byte

	randomQuote := rand.Intn(len(x.quotacoes))
	w.Header().Set("Content-Type", "application/json")
	fmt.Printf("Returning: %v\n", x.quotacoes[randomQuote])

	jsonResponse, err := json.Marshal(x.quotacoes[randomQuote])

	if err != nil {
		log.Printf("Error Marshalling: %v", err)
	}

	w.Write(jsonResponse)
}

func (x *quotes) load(fname string) error {
	records, err := readData(fname)

	if err != nil {
		return err
	}

	for _, r := range records {
		record := singleQuote{
			Quote:  r[0],
			Author: r[1],
		}

		x.quotacoes = append(x.quotacoes, record)

		fmt.Printf("%v\n", record)
		fmt.Printf(("=>==========================================================================================\n"))

	}

	return nil
}

func main() {

	myquotes := &quotes{}

	if err := myquotes.load("quotes.csv"); err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/", myquotes.Handler).Methods("GET", "OPTIONS")

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
