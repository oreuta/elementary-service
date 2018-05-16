package sequence

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/oreuta/elementary-service/src/models/sequence"
	"regexp"
	"io/ioutil"
)

// `{"length": 5, "minSq": 10}`
type inputData struct {
	Length int `json:"length"`
	MinSq int  `json:"minSq"`
}

// `{"sequence_of_natural_digits":[1, 4, 7]}`
type outputData struct {
	Sequence string `json:"sequence_of_natural_digits"`
}

const serviceName = "Sequence"

func logError(err error) {
	log.Printf("%s: ERROR %q", serviceName, err.Error())
}

var re = regexp.MustCompile("^[0-9]+$")

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
	input := inputData{}
	err = json.Unmarshal(body, &input)
	if err != nil {
		logError(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}


	arr, err := sequence.GetSquares(input.Length, input.MinSq)
	if err != nil {
		logError(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Printf("%s: output data %v", serviceName, arr)

	str := sequence.PrintWithCommas(arr)

	output := outputData{
		Sequence: str,
	}

	outputJSON, err := json.Marshal(output)
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
