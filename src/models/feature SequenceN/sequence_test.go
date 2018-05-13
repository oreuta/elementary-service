package main

import (
	"testing"
)

func TestForValidationofInputData(t *testing.T) {
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
			gotlenght, gotminSq, gotError := Checkthedata(testCase.inputlenght, testCase.inputminSq)
			if testCase.needError != (gotError !=nil) {
				t.Errorf("For the returned error: want %v but got %v", gotlenght, gotminSq)
			}
		})
	}

}

func TestForSquares(t *testing.T) {
	var testCases1 = []struct {
		name           string
		lenght1        int
		minSq1         int
		expectedResult []int
		needError      bool
	}{

		{"Test for getting the square for lenght=1, minSq=2",
			1,
			2,
			[]int{4},
			false},

		{"Test for getting the square for lenght=4, minSq=5",
			4,
			5,
			[]int{9, 16, 25, 36},
			false},

		{"Test for getting the square for lenght=4, minSq=0",
			4,
			0,
			[]int{1, 4, 9, 16},
			false},

		{"Negative test for wrong expected results/the square for lenght=0, minSq=3",
			0,
			3,
			[]int{4, 9, 16},
			true},

		{"Test for getting the square of natural numbers for lenght=8, minSq=-10",
			8,
			-10,
			[]int{1, 4, 9, 16, 25, 36, 49, 64},
			true},
	}
	for _, testCase1 := range testCases1 {
		testCase1 := testCase1
		t.Run(testCase1.name, func(t *testing.T) {
			actualResult, err := GetSquares(testCase1.lenght1, testCase1.minSq1)
			if testCase1.needError != (err != nil) {
				t.Errorf("For returned error: want %v but got %v", actualResult, err)
			}
			})

	}
}

