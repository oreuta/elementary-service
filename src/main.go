package main

import (
	"log"
	"net/http"

	"services/squarert"
	"services/sequence"
)

func main() {
	log.Println("Elementary-service started...")
	http.HandleFunc("/squarert", squarert.Handler)
	http.HandleFunc("/sequence", sequence.Handler)
	log.Fatal(http.ListenAndServe(":9090", nil))



}
