package squarert

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/elementary-service/src/models/squarert"
)

// `{"numbers":[1,2,3]}`
type inputNumbers struct {
	Numbers []int
}

// `{"square_roots":[1,1.41,1.7]}`
type outputNumbers struct {
	SquareRoots []float64 `json:"square_roots"`
}

const serviceName = "SquareRoots"

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
	numbers := inputNumbers{}
	err = json.Unmarshal(body, &numbers)
	if err != nil {
		logError(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	outputData, err := squarert.SquareRoot(numbers.Numbers)
	if err != nil {
		logError(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Printf("%s: output data %v", serviceName, outputData)

	outputStruct := outputNumbers{
		SquareRoots: outputData,
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
