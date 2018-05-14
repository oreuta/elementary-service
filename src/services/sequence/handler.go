package sequence

import (
	"encoding/json"
	"log"
	"net/http"
	"models/sequence"
	"regexp"
)

// `{"numbers":[1,2,3]}`
type inputData struct {
	lenght string
	minSq string
}

// `{"sequence_of_natural_digits":[1, 4, 7]}`
type outputData struct {
	Sequence string `json:"sequence_of_natural_digits"`
}

const serviceName = "SquareRoots"

func logError(err error) {
	log.Printf("%s: ERROR %q", serviceName, err.Error())
}
r := regexp.MustCompile("^[0-9]+$")

// Handler is a REST wrapper for SquareRoot function
func Handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s: start", serviceName)
	/*defer log.Printf("%s: stop", serviceName)*/

	/*body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()*/

	/*log.Printf("%s: input data %s", serviceName, body)
	numbers := inputData{}
	err = json.Unmarshal(body, &numbers)
	if err != nil {
		logError(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}*/

	ll1, ms1, err := sequence.Validatetheinput("4", "5", r)
	if err != nil {
		logError(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Printf("%d, %d: output data %v", serviceName, ll1, ms1)

	arr, err := sequence.GetSquares(ll1, ms1)
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
