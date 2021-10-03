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
	//lang := 0
	quoteID := ""

	db, err := sql.Open("sqlite3", "./quotesqlite")
	checkError("Error opening database: ", err)

	rows, err := db.Query("SELECT quotes.quote_id, quotes.quote, authors.author_name FROM quotes INNER JOIN authors ON quotes.author_id = authors.author_id ORDER BY RANDOM() LIMIT 1")
	checkError("Error executing query: ", err)

	for rows.Next() {
		err = rows.Scan(&quoteID, &myQuote.Quote, &myQuote.Author)
	}
	checkError("Error, no data returned from query: ", err)

	rows.Close()

	db.Close()
	return (myQuote)

}

// Get quote handler
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

// Create quote handler
func CreateHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "./quotesqlite")
	checkError("Error opening database: ", err)
	defer db.Close()

	// Processes parameters received through http.Request
	author := r.FormValue("author")
	quote := r.FormValue("quote")
	log.Printf("Created: %v => %v\n ", author, quote)

	// Saves new quote to database
	stmt, err := db.Prepare(`
		INSERT INTO quotes(idiom, quote, author)
		VALUES(?, ?, ?)
	`)
	checkError("Prepare query error: ", err)

	_, err = stmt.Exec(1, quote, author)
	checkError("Execute query error: ", err)
}

// Update quote handler
func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Updated: %v\n", "sample author,sample quote")
}

// Delete quote handler
func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Deleted: %v\n", "sample author,sample quote")
}

func main() {
	// checks if database file exists and aborts app if it does not
	checkIfDabaseExists()

	// defines routes
	router := mux.NewRouter()
	router.HandleFunc("/", Handler).Methods("GET", "OPTIONS")
	router.HandleFunc("/create", CreateHandler).Methods("POST", "OPTIONS")
	router.HandleFunc("/update", UpdateHandler).Methods("PUT", "OPTIONS")
	router.HandleFunc("/delete", DeleteHandler).Methods("DELETE", "OPTIONS")

	//sets listener on port 8080 unless PORT environent variable is defined
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
