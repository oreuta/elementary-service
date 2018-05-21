package trianglesSort

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"github.com/oreuta/elementary-service/src/models/trianglesSort"
	"bytes"
)

func TestHandler(t *testing.T) {

	testCases := []struct {
		name            string
		inputBody       string
		expResponseBody string
		needError       bool
	}{
		{
			name: "no err, regular request",
			inputBody:`{"triangles":[{"vertices": "ABC","a": 10,"b": 20,"c": 22.36},{"vertices":"CBA","a":16,"b":15,"c": 6}]}`,
			expResponseBody:`{"triangles":[{"vertices":"ABC","a":10,"b":20,"c":22.36,"square":0},{"vertices":"CBA","a":16,"b":15,"c":6,"square":0}]}`,
			needError:false,
		},
		{
			name: "need err, bad request, want number but got string",
			inputBody:`{"triangles":[{"vertices": "ABCsd","a": "10","b": 20,"c": 22.36},{"vertices":"CBA","a":16,"b":15,"c": 6}]}`,
			expResponseBody:``,
			needError:true,
		},
		{
			name: "need err, bad request, empty",
			inputBody:`{}`,
			expResponseBody:``,
			needError:true,
		},
	}

	//Mock
	trianglesSquare = func(trianglesToSortSlice []trianglesSort.Triangle) ([]trianglesSort.Triangle, error) {
		return trianglesToSortSlice,nil
	}


	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			req, err := http.NewRequest("POST", "/trianglesSort",bytes.NewReader([]byte(testCase.inputBody)))
			if err != nil {
				t.Fatal(err)
			}
			rr := httptest.NewRecorder()

			handler := http.HandlerFunc(Handler)
			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != http.StatusOK  && !testCase.needError{
				t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
			}

			temp :=rr.Body.String()
			if temp != testCase.expResponseBody && !testCase.needError{
				t.Errorf("handler returned unexpected body: %v",temp)
			}
		})

	}

}
