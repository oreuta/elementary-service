package main

import (
	"log"
	"net/http"
	"github.com/oreuta/elementary-service/src/services/palindrome"
	"github.com/oreuta/elementary-service/src/services/squarert"
	"github.com/oreuta/elementary-service/src/services/fibo"
	"github.com/oreuta/elementary-service/src/services/sequence"
)

func main() {

	log.Println("Elementary-service started...")
	http.HandleFunc("/fibo", fibo.Handler)
	http.HandleFunc("/subpalindromes", palindrome.Handler)
	http.HandleFunc("/squarert", squarert.Handler)
	http.HandleFunc("/sequence", sequence.Handler)
	log.Fatal(http.ListenAndServe(":9090", nil))
}



