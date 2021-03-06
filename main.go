package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var (
	db            *sql.DB
	err           error
	authenticated = false
)

func main() {
	db, err = sql.Open("sqlite3", "db.sqlite3")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	// test connection
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	os.Setenv("PORT", "2301")
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	// route
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/logout", logoutHandler)
	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/forgot", forgotHandler)
	http.HandleFunc("/listAdmin", listAdminHandler)
	http.HandleFunc("/list", listHandler)
	http.HandleFunc("/createRetailer", createRetailerHandler)
	http.HandleFunc("/createCustomer", createCustomerHandler)
	http.HandleFunc("/updateRetailer", updateRetailerHandler)
	http.HandleFunc("/updateCustomer", updateCustomerHandler)
	http.HandleFunc("/deleteCustomer", deleteCustomerHandler)
	http.HandleFunc("/deleteRetailer", deleteRetailerHandler)
	http.HandleFunc("/generatePdf", generatePDFHandler)
	http.Handle("/css/",
		http.StripPrefix("/css/", http.FileServer(http.Dir("./css"))),
	)
	http.ListenAndServe(":"+port, nil)
}
