package sequence

import (
	"testing"
)

func TestsForValidationofInputData(t *testing.T) {
	var testCases = []struct {
		name        string
		inputlenght string
		inputminSq  string
		needError   bool
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
			if gotlenght, gotminSq, gotError := Checkthedata(testCase.inputlenght, testCase.inputminSq); testCase.needError != (gotError != nil) {
				t.Errorf("for returned error: want %f but got %v", testCase.name, gotlenght, gotminSq, gotError)
			}
		})
	}

}

func TestsForSquares(t *testing.T) {
	var testCases1 = []struct {
		name           string
		lenght1        int
		minSq1         int
		expectedResult []int
		needError      bool
	}{
		{"Test for dealing with 0",
			0,
			0,
			[]int{0},
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
	for _, testCase1 := range testCases1 {
		testCase1 := testCase1
		t.Run(testCase1.name, func(t *testing.T) {
			actualResult, err := GetSquares(testCase1.lenght1, testCase1.minSq1)
			if testCase1.needError != (err != nil) { //XOR
				t.Errorf("For returned error: want <nil> but got %v", err)
			}
			for i := range testCase1.expectedResult {
				if testCase1.expectedResult[i] != actualResult[i] {
					t.Errorf("want %f but got %f", testCase1.expectedResult[i], actualResult[i])
				}
			}
		})

	}
}

func TestsForPrintingwithCommas(t *testing.T) {
	var testCases2 = []struct {
		name           string
		inputdata      []int
		expectedResult string
		needError      bool
	}{
		{"Test for printing without commas",
			[]int{1, 4, 9, 16},
			"1 4 9 16",
			true},

		{"Test for printing with commas",
			[]int{1, 4, 9, 16},
			"1,4,9,16",
			false},
	}

	for _, testCase2 := range testCases2 {
		testCase2 := testCase2
		t.Run(testCase2.name, func(t *testing.T) {
			actualResult := Printwithcommas(testCase2.inputdata)
			if testCase2.needError != true {
				t.Errorf("For returned error: want <nil> but got %v", testCase2.needError)
			}
			for i := range testCase2.expectedResult {
				if testCase2.expectedResult[i] != actualResult[i] {
					t.Errorf("want %f but got %f", testCase2.expectedResult[i], actualResult[i])
				}
			}
		})

	}
}
