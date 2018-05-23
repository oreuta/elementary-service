package services

import (
	"net/http"
	"log"
	"github.com/oreuta/elementary-service/src/services/squarert"
	"github.com/oreuta/elementary-service/src/services/trianglesSort"
	"github.com/oreuta/elementary-service/src/services/fibo"
	"github.com/oreuta/elementary-service/src/services/palindrome"
	"github.com/oreuta/elementary-service/src/services/sequence"
	"time"
)

//Router is a function that starts server and routes http requests
func Router() {
	log.Println("Elementary-service started...")

	trianglesSortHandler := http.HandlerFunc(trianglesSort.Handler)

	http.HandleFunc("/fibo", fibo.Handler)
	http.HandleFunc("/subpalindromes", palindrome.Handler)
	http.HandleFunc("/squarert", squarert.Handler)
	http.HandleFunc("/sequence", sequence.Handler)
	http.Handle("/trianglesSort", handlerLogWrapper(trianglesSortHandler))
	http.HandleFunc("/luckyTickets", trianglesSort.Handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handlerLogWrapper(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Executing trianglesSquareSort handler")
		startTime := time.Now()
		next.ServeHTTP(w, r)
		endTime:=time.Now()
		log.Printf("Executing of the trianglesSquareSort handler has finished, exec time: %v\n", endTime.Sub(startTime))
	})
}
