package services

import (
	"net/http"
	"log"
	"services/squarert"
	"services/trianglesSort"
	"services/fibo"
	"services/palindrome"
	"services/sequence"
	"os"
	"services/luckyTickets"
)

//Router is a function that starts server and routes http requests
func Router() {
	port := os.Getenv("PORT")
	log.Println("Elementary-service started...")

	http.HandleFunc("/fibo", fibo.Handler)
	http.HandleFunc("/subpalindromes", palindrome.Handler)
	http.HandleFunc("/squarert", squarert.Handler)
	http.HandleFunc("/sequence", sequence.Handler)
	http.HandleFunc("/trianglesSort", trianglesSort.Handler)
	http.HandleFunc("/luckyTickets", luckyTickets.Handler)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
