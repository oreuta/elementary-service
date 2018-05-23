package services

import (
	"net/http"
	"log"
	"github.com/oreuta/elementary-service/src/services/squarert"
	"github.com/oreuta/elementary-service/src/services/trianglesSort"
	"github.com/oreuta/elementary-service/src/services/palindrome"
	"github.com/oreuta/elementary-service/src/services/sequence"
	"github.com/oreuta/elementary-service/src/services/fibo"
)

//Router is a function that starts server and routes http requests
func Router() {
	log.Println("Elementary-service started...")

	http.HandleFunc("/fibo", fibo.Handler)
	http.HandleFunc("/subpalindromes", palindrome.Handler)
	http.Handle("/LogSubpalindromes", palindrome.WrappedLogHandler(http.HandlerFunc(palindrome.Handler)))
	http.HandleFunc("/squarert", squarert.Handler)
	http.HandleFunc("/sequence", sequence.Handler)
	http.HandleFunc("/trianglesSort", trianglesSort.Handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
