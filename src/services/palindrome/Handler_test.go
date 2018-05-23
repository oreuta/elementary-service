package palindrome

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"bytes"
	"github.com/oreuta/elementary-service/src/models/palindrome"
)

func TestHandler(t *testing.T) {

	testCases := []struct {
		name         string
		inputBody    string
		expectedBody string
		needError    bool
	}{
		{
			name:         "no err, input string 'kayak'",
			inputBody:    `{"InputString": "kayak"}`,
			expectedBody: `{"sub_palindromes":["kayak"]}`,
			needError:    false,
		},
		{
			name:      "err, input string has 1 symbol",
			inputBody: `{"InputString": "k"}`,
			needError: true,
		},
		{
			name:      "err, input string is empty",
			inputBody: `{"InputString": ""}`,
			needError: true,
		},
	}


	palindrome.TestMockHandler()

	for _, testCase := range testCases {
		testCase := testCase
		body := bytes.NewReader([]byte(testCase.inputBody))
		t.Run(testCase.name, func(t *testing.T) {
			req, _ := http.NewRequest("POST", "/subpalindromes", body)

			w := httptest.NewRecorder()
			Handler(w, req)

			if w.Code != http.StatusOK && !testCase.needError {
				t.Errorf("handler got %v, but want %v", w.Code, http.StatusOK)
			}

			if w.Body.String() != testCase.expectedBody && !testCase.needError {
				t.Errorf("handler returned unexpected body: %v", w.Body.String())
			}
		})

	}

}
