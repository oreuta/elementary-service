package trianglesSort

import (
	"github.com/oreuta/elementary-service/src/models/trianglesSort"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"log"
)

/*TrianglesBody Input json example
{
	"triangles": [{
		"vertices": "ABC",
		"a": 10,
		"b": 20,
		"c": 22.36
	}, {
		"vertices": "CBA",
		"a": 16,
		"b": 15,
		"c": 6
	}, {
		"vertices": "BAC",
		"a": 10,
		"b": 9,
		"c": 16
	}]
}
 */
type TrianglesBody struct {
	Triangles []trianglesSort.Triangle `json:"triangles"`
}

const serviceName = "TriangleSort"

func logError(err error) {
	log.Printf("%s: ERROR %q", serviceName, err.Error())
}

// Handler is a REST wrapper for TrianglesSort function
func Handler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	trianglesToSort := TrianglesBody{}
	err = json.Unmarshal(body, &trianglesToSort)
	if err != nil {
		logError(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	outputData, err := trianglesSort.TrianglesSquareSort(trianglesToSort.Triangles)
	if err != nil {
		logError(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	outputStruct := TrianglesBody{
		Triangles: outputData,
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
		return
	}
}