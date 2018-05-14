package services

import (
	"net/http"
	"log"
	"github.com/oreuta/elementary-service/src/services/squarert"
	"github.com/oreuta/elementary-service/src/services/trianglesSort"
)

//Router is a function that starts server and routes http requests
func Router() {
	http.HandleFunc("/squarert", squarert.Handler)
	http.HandleFunc("/trianglesSort", trianglesSort.Handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
