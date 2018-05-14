package main

import (
	"log"
	"net/http"
	"github.com/oreuta/elementary-service/src/services/palindrome"
	"github.com/oreuta/elementary-service/src/services/squarert"
	"github.com/oreuta/elementary-service/src/services/fibo"
)

func main() {

	log.Println("Elementary-service started...")
	http.HandleFunc("/fibo", fibo.Handler)
	http.HandleFunc("/subpalindromes", palindrome.Handler)
	http.HandleFunc("/squarert", squarert.Handler)
	log.Fatal(http.ListenAndServe(":9090", nil))
}
