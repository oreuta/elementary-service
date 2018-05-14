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
				t.Errorf("for returned error: want %s but got %d, %d, %s", testCase.name, gotlenght, gotminSq, gotError)
			}
		})
	}
}


func TestSquares(t *testing.T) {
	var testCases = []struct {
		name           string
		lenght1        int
		minSq1         int
		expectedResult []int
		needError      bool
	}{
		{"Test for dealing with 0",
			0,
			0,
			[]int{},
			true},

		{"Test for getting the square for lenght=1, minSq=2",
			1,
			2,
			[]int{4},
			true},

		{"Test for getting the square for lenght=4, minSq=5",
			4,
			5,
			[]int{9, 16, 25, 36},
			true},

		{"Test for getting the square for lenght=4, minSq=5",
			4,
			0,
			[]int{1, 4, 9, 16},
			true},
	}
	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			actualResult, err := GetSquares(testCase.lenght1, testCase.minSq1)
			if testCase.needError == (err != nil) {
				t.Errorf("For returned error: want <nil> but got %v", err)
			}
			for i,_ := range testCase.expectedResult {
				if testCase.expectedResult[i] != actualResult[i] {
					t.Errorf("want %d but got %d", testCase.expectedResult[i], actualResult[i])
				}
			}
		})
	}
}

func TestPrintingwithCommas(t *testing.T) {
	var testCases = []struct {
		name           string
		inputdata      []int
		expectedResult string
		needError      bool
	}{
		{"Test for printing with commas",
			[]int{1, 4, 9, 16},
			"1, 4, 9, 16",
			false,
	},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			actualResult := PrintWithCommas(testCase.inputdata)
			if testCase.expectedResult != actualResult {
				t.Error("Print with commas return incorrect data")
			}
		})

	}
}
