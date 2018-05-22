package luckyTickets

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"bytes"
)


func TestHandler(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	//body := bytes.NewReader([]byte("{\"min\": 100000, \"max\": 10003033}"))

	bodies := []struct{
		expectedStatus int
		inputBody []byte
	}{
		{
			expectedStatus: 200,
			inputBody: []byte("{\"min\": 100000, \"max\": 100030}"),
		},
		{
			expectedStatus: 400,
			inputBody: []byte("{\"min\": 100000, \"max\": \"wqeqwe\"}"),
		},
	}

	for _, body := range bodies{
		body := body
		req, err := http.NewRequest("POST", "/luckyTickets", bytes.NewReader(body.inputBody))
		if err != nil {
			t.Fatal(err)
		}

		// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(Handler)

		// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
		// directly and pass in our Request and ResponseRecorder.
		handler.ServeHTTP(rr, req)

		// Check the status code is what we expect.
		if status := rr.Code; status != body.expectedStatus{
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, body.expectedStatus)
		}
	}

	// Check the response body is what we expect.
	//expectedOutputData := `{"Method":0,"CountFirstMethod":0,"CountSecondMethod":0}`



	//if rr.Body.String() != expectedOutputData {
	//	t.Errorf("handler returned unexpected body: got %v want %v",
	//		rr.Body.String(), expectedOutputData)
	//}
}