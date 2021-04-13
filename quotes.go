package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

type singleQuote struct {
	Quote  string
	Author string
}

func Handler(w http.ResponseWriter, r *http.Request) {
	var jsonResponse []byte

	myQuote := singleQuote{}
	lang := 0

	//retrieves random quote from table quotes on sqlite database.

	db, err := sql.Open("sqlite3", "./quotesqlite")
	if err != nil {
		log.Printf("Error opening database: %v", err)
	}

	rows, err := db.Query("SELECT * FROM quotes ORDER BY RANDOM() LIMIT 1")
	if err != nil {
		log.Printf("Error executing query: %v", err)
	}

	for rows.Next() {
		err = rows.Scan(&lang, &myQuote.Quote, &myQuote.Author)
	}

	//todo  - add err check here

	rows.Close()

	db.Close()

	// sets headers, marshalls and returns quote + author JSON
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	log.Printf("Returning: %v\n", myQuote)

	jsonResponse, err = json.Marshal(myQuote)

	if err != nil {
		log.Printf("Error Marshalling: %v", err)
	}

	w.Write(jsonResponse)
}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/", Handler).Methods("GET", "OPTIONS")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Assuming port# %s as standard", port)
	} else {
		log.Printf("Listening on port# %s", port)
	}

	corsObj := handlers.AllowedOrigins([]string{"*"})

	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS(corsObj)(router)))
}
