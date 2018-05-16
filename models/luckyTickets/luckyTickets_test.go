package luckyTickets

import (
	"testing"
	"reflect"
)

func TestGetLuckyTickets(t *testing.T){
	testCases := []struct{
		name string
		needError bool
		inputData TaskContext
		expected Winner
	}{
		{
			name: "no error",
			needError: false,
			inputData: TaskContext{100000, 100010},
			expected: Winner {1, 2, 0},
		},
		{
			name: "numbers must be positive",
			needError: true,
			inputData: TaskContext{-999999, 100000},
			expected: Winner {0, 0, 0},
		},
		{
			name: "min must be lower than max and their length must be 6",
			needError: true,
			inputData: TaskContext{75, 50},
			expected: Winner {0, 0, 0},
		},
		{
			name: "their length must be 6",
			needError: true,
			inputData: TaskContext{1234567, 12345678},
			expected: Winner {0, 0, 0},
		},
	}

	for _, testCase := range testCases{
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			gotResult, err := GetLuckyTickets(&testCase.inputData)
			if testCase.needError != (err != nil){
				t.Errorf("for returned error: want <nil> but got %v", err)
			}
			if !reflect.DeepEqual(gotResult, testCase.expected){
				t.Fatalf("\n expected struct: %+v, \n but got: %+v", testCase.expected, gotResult)
			}
		})
	}



}