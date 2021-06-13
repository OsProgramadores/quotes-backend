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

/*
   Quote struture definition
   Quote  == quote text
   Author == quote author
*/

type singleQuote struct {
	Quote  string
	Author string
}

// Checks for fatal error and displays error message and abort program if applicable
func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}

// Checks if sqlite database file exists and aborts program if not
func checkIfDabaseExists() {
	log.Printf("Checking if Database exists...")
	_, err := os.Stat("./quotesqlite")
	if os.IsNotExist(err) {
		checkError("Database not found: ", err)
	}
	log.Printf("Database ok.")
}

// Retrieves random quote from sqlite database table quotes
func readRandomQuote() singleQuote {
	myQuote := singleQuote{}
	lang := 0

	db, err := sql.Open("sqlite3", "./quotesqlite")
	checkError("Error opening database: ", err)

	rows, err := db.Query("SELECT * FROM quotes ORDER BY RANDOM() LIMIT 1")
	checkError("Error executing query: ", err)

	for rows.Next() {
		err = rows.Scan(&lang, &myQuote.Quote, &myQuote.Author)
	}
	checkError("Error, no data returned from query: ", err)

	rows.Close()

	db.Close()
	return (myQuote)

}

// Http handler
func Handler(w http.ResponseWriter, r *http.Request) {
	// retrieves random quote from sqlite database table quotes
	myQuote := readRandomQuote()

	// sets headers, marshalls and returns quote + author JSON
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	jsonResponse, err := json.Marshal(myQuote)
	checkError("Error Marshalling: ", err)

	w.Write(jsonResponse)
	log.Printf("Returned: %v\n", myQuote)
}

func main() {
	// checks if database file exists and aborts app if it does not
	checkIfDabaseExists()

	// defines http handler and sets listener on port 8080 unless PORT environent variable is defined
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
