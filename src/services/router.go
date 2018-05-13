package services

import (
	"net/http"
	"log"
	"services/squarert"
	"services/trianglesSort"
)

//Router is a function that starts server and routes http requests
func Router() {
	http.HandleFunc("/squarert", squarert.Handler)
	http.HandleFunc("/trianglesSort", trianglesSort.Handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
