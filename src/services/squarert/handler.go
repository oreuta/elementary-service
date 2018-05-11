package squarert

import (
	"net/http"
	"encoding/json"
	"elementary-service/src/models/squarert"
	"log"
)

// `{"numbers":[1,2,3]}`
type inputNumbers struct {
	Numbers []int
}

// `{"square_roots":[1,1.41,1.7]}`
type outputNumbers struct {
	SquareRoots []float64 `json:"square_roots"`
}

// Handler is a handler for SquareRoot
func Handler(w http.ResponseWriter, r *http.Request) {
	log.Println("handler started...")
	body, err := r.GetBody()
	log.Printf("body=%v err=%v\n", body, err)
	if err != nil || body == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer body.Close()

	data := make([]byte, 0)

	_, err = body.Read(data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Printf("data=%v\n", data)
	numbers := inputNumbers{}

	err = json.Unmarshal(data, &numbers)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	output, err := squarert.SquareRoot(numbers.Numbers)
	log.Printf("output=%v err=%v\n", output, err)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	outputJSON, err := json.Marshal(output)
	log.Printf("json=%v err=%v\n", outputJSON, err)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = w.Write(outputJSON)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	header := w.Header()
	header.Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}