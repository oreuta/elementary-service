package sequence

import (
	"testing"
)

func TestValidationInputData(t *testing.T) {

	var testCases = []struct {
		name            string
		inputLenght     int
		inputMinSquare  int
		needError       bool
	}{
		{"Test for two positive numbers",
			123,
			34,
			false,
		},

		{"Test for two negative numbers",
			-123,
			-24,
			true,
		},

		{"Test for two zero numbers",
			0,
			0,
			true,
		},

		{"Test for negative input lenght",
			-123,
			24,
			true,
		},

		{"Test for negative square",
			123,
			-24,
			true,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			if gotlenght, gotminSq, gotError := ValidateTheInput(testCase.inputLenght, testCase.inputMinSquare); testCase.needError != (gotError != nil) {
				t.Errorf("for returned error: want %s but got %d, %d, %s", testCase.name, gotlenght, gotminSq, gotError)
			}
		})
	}
}

func TestSquares(t *testing.T) {
	var testCases = []struct {
		name           string
		inputLenght     int
		inputMinSquare  int
		expectedResult []int
		needError      bool
	}{
		{"Test for dealing with 0",
			0,
			0,
			[]int{},
			false,
		},

		{"Test for getting the square for lenght=1, minSq=2",
			1,
			2,
			[]int{4},
			false,
		},

		{"Test for getting the square for lenght=4, minSq=5",
			4,
			5,
			[]int{9, 16, 25, 36},
			false,
		},
	}
	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			actualResult, err := GetSquares(testCase.inputLenght, testCase.inputMinSquare)
			if err != nil {
				t.Errorf("For returned error: want <nil> but got %v", err)
			}
			for i := range testCase.expectedResult {
				if testCase.expectedResult[i] != actualResult[i] {
					t.Errorf("want %d but got %d", testCase.expectedResult[i], actualResult[i])
				}
			}
		})
	}
}

func TestPrintWithCommas(t *testing.T) {
	var testCases = []struct {
		name           string
		inputData      []int
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
			actualResult := PrintWithCommas(testCase.inputData)
			if testCase.expectedResult != actualResult {
				t.Error("Print with commas return incorrect data")
			}
		})

	}
}