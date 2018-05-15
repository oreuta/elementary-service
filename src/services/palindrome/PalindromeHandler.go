package palindrome
import (
	"encoding/json"
	"log"
	"net/http"
	"io/ioutil"
	palindrome2 "github.com/oreuta/elementary-service/src/models/palindrome"
)


type inputString struct {
	InputString string
}


type outputStrings struct {
	SubPalindromes string `json:"sub_palindromes"`
}

const serviceName = "SubPalindromes"

func logError(err error) {
	log.Printf("%s: ERROR %q", serviceName, err.Error())
}

// Handler is a REST wrapper for SubPalindromes function
func Handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s: start", serviceName)
	defer log.Printf("%s: stop", serviceName)

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

	outputData, err := palindrome2.FindSubPalindromes(strings.InputString)
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
