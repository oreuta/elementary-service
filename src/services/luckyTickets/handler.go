package luckyTickets

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/oreuta/elementary-service/src/models/luckyTickets"
	"time"
)

type Winner struct {
	method                              int
	countFirstMethod, countSecondMethod int
}

const serviceName = "LuckyTickets"

func logError(err error) {
	log.Printf("%s: ERROR %q", serviceName, err.Error())
}

func LogHandler(h http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\r\n Executing luckyTicketsLogHandler... start-time: %v \r\n", time.Now().Nanosecond() )
			Handler(w, r)
		log.Printf("\r\n LuckyTicketsLogHandler has finished. end-time: %v \r\n", time.Now().Nanosecond())
	})
}

// Handler is a REST wrapper for LuckyTickets-function
func Handler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	log.Printf("\r\n %s: input data %s", serviceName, body)

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
	log.Printf("\r\n %s: output data %v", serviceName, Winner)

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
