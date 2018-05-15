package services

import (
	"net/http"
	"log"
	"github.com/oreuta/elementary-service/src/services/squarert"
	"github.com/oreuta/elementary-service/src/services/trianglesSort"
	"github.com/oreuta/elementary-service/src/services/fibo"
	"github.com/oreuta/elementary-service/src/services/palindrome"
	"github.com/oreuta/elementary-service/src/services/sequence"
	"os"
)

//Router is a function that starts server and routes http requests
func Router() {
	log.Println("Elementary-service started...")
	port := os.Getenv("PORT")
	http.HandleFunc("/fibo", fibo.Handler)
	http.HandleFunc("/subpalindromes", palindrome.Handler)
	http.HandleFunc("/squarert", squarert.Handler)
	http.HandleFunc("/sequence", sequence.Handler)
	http.HandleFunc("/trianglesSort", trianglesSort.Handler)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
