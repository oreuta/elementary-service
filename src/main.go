package main

import (
	"log"
	"net/http"

	"github.com/oreuta/elementary-service/src/services/squarert"
)

func main() {
	log.Println("Elementary-service started...")
	http.HandleFunc("/squarert", squarert.Handler)
	log.Fatal(http.ListenAndServe(":9090", nil))
}
