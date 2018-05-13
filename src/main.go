package main

import (
	"log"
	"net/http"
	"src/services/palindrome"

)

func main() {
	log.Println("Elementary-service started...")
	http.HandleFunc("/squarert", palindrome.Handler)
	log.Fatal(http.ListenAndServe(":9090", nil))
}
