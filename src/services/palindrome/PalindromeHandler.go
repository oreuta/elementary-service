package palindrome
import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"src/models/palindrome"

)


type inputString struct {
	Strings string
}


type outputStrings struct {
	HasPalindrome []string `json:"has_palindrome"`
}

const serviceName = "HasPalindrome"

func logError(err error) {
	log.Printf("%s: ERROR %q", serviceName, err.Error())
}

// Handler is a REST wrapper for SquareRoot function
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

	outputData, err := palindrome.HasPalindrome(strings.Strings)
	if err != nil {
		logError(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Printf("%s: output data %v", serviceName, outputData)

	outputStruct := outputStrings{
		HasPalindrome: outputData,
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
