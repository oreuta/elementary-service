package sequence

import (
	"testing"
	"regexp"
)

func TestValidationofInputData(t *testing.T) {
	r := regexp.MustCompile("^[0-9]+$")
	var testCases = []struct {
		name            string
		inputlenght     string
		inputminsquare  string
		needError         bool
	}{
		{"Test for two positive numbers",
			"123",
			"24",
			false},

		{"Test for two negative numbers",
			"-123",
			"-24",
			true},

		{"Test for two zero numbers",
			"0",
			"0",
			true},

		{"Test for positive number and word",
			"123",
			"hello",
			true},

		{"Test for two empty strings",
			" ",
			" ",
			true},

		{"Test for two numbers with separators",
			"123.56",
			"24.51",
			true},

		{"Test for two words",
			"love",
			"été",
			true},

		{"Test for two dots",
			".",
			".",
			true},

		{"Test for zero and one",
			"1",
			"0",
			true},

		{"Test for one negative and one positive number",
			"-123",
			"24",
			true},
	}


	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			if gotlenght, gotminSq, gotError := Validatetheinput(testCase.inputlenght, testCase.inputminsquare, r); testCase.needError != (gotError != nil) {
				t.Errorf("for returned error: want %f but got %v", testCase.name, gotlenght, gotminSq, gotError)
			}
		})
	}
}

