package services

import (
	"net/http"
	"log"
	"services/squarert"
	"services/trianglesSort"
	"services/fibo"
	"services/palindrome"
	"services/sequence"
)

//Router is a function that starts server and routes http requests
func Router() {
	log.Println("Elementary-service started...")

	http.HandleFunc("/fibo", fibo.Handler)
	http.HandleFunc("/subpalindromes", palindrome.Handler)
	http.HandleFunc("/squarert", squarert.Handler)
	http.HandleFunc("/sequence", sequence.Handler)
	http.HandleFunc("/trianglesSort", trianglesSort.Handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
