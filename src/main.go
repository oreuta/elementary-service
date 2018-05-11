package main

import (
	"net/http"
	"elementary-service/src/services/squarert"
	"log"
)

func main() {
	log.Println("started...")
	http.HandleFunc("/sqarert", squarert.Handler)
	log.Fatal(http.ListenAndServe(":9090", nil))
}


