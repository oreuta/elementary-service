package palindrome

import (
	"encoding/json"
	"log"
	"net/http"
	"io/ioutil"
	"github.com/oreuta/elementary-service/src/models/palindrome"
	"time"
)

type inputString struct {
	InputString string
}

type outputStrings struct {
	SubPalindromes []string `json:"sub_palindromes"`
}

const serviceName = "SubPalindromes"

func logError(err error) {
	log.Printf("%s: ERROR %q", serviceName, err.Error())
}

// Handler is a REST wrapper for SubPalindromes function
func Handler(w http.ResponseWriter, r *http.Request) {
	//log.Printf("%s: start", serviceName)
	//defer log.Printf("%s: stop", serviceName)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	log.Printf("%s: input data %s", serviceName, body)
	strings := inputString{}
	err = json.Unmarshal(body, &strings)
	if err != nil {
		logError(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	outputData, err := palindrome.FindSubPalindromes(strings.InputString)
	if err != nil {
		logError(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Printf("%s: output data %v", serviceName, outputData)

	outputStruct := outputStrings{
		SubPalindromes: outputData,
	}

	outputJSON, err := json.Marshal(outputStruct)
	if err != nil {
		logError(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(outputJSON)
	if err != nil {
		logError(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func WrappedLogHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		log.Printf("Start time of palindrome handler is %s s", time.Now().Round(time.Nanosecond).Format("15:04:05.999999999"))
		Handler(w, r)
		log.Printf("Palindrome handler has finished his work at %s s", time.Now().Round(time.Nanosecond).Format("15:04:05.999999999"))
	})
}
