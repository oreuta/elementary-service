package luckyTickets

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"models/luckyTickets"
)

type Winner struct {
	method                              int
	countFirstMethod, countSecondMethod int
}

const serviceName = "LuckyTickets"

func logError(err error) {
	log.Printf("%s: ERROR %q", serviceName, err.Error())
}

// Handler is a REST wrapper for LuckyTickets-function
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
	numbers := new(luckyTickets.TaskContext)
	err = json.Unmarshal(body, &numbers)
	if err != nil {
		logError(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	Winner, err := luckyTickets.GetLuckyTickets(numbers)
	if err != nil {
		logError(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Printf("%s: output data %v", serviceName, Winner)

	outputJSON, err := json.Marshal(Winner)
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
