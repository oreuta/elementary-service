package main

import (
	/*"log"
	"net/http"

	"services/squarert"*/
	"models/sequence"
	"regexp"
	"fmt"
)

func main() {
	/*log.Println("Elementary-service started...")
	http.HandleFunc("/squarert", squarert.Handler)
	http.HandleFunc("/sequence", sequence.Handler)
	log.Fatal(http.ListenAndServe(":9090", nil))*/

	r := regexp.MustCompile("^[0-9]+$")
	ll1, ms1, err := sequence.Validatetheinput("4", "0", r)
	arr, err := sequence.GetSquares(ll1, ms1)
	str := sequence.PrintWithCommas(arr)
	fmt.Println("Numerical sequence separated by commas:  ", str)
	fmt.Println(err)
}



