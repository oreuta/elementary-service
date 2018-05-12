package luckyTickets

import (
	"testing"
	"reflect"
	"fmt"
)

func TestGetLuckyTickets(t *testing.T){
	testObjects := []struct{
		name string
		needError bool
		inputData context
		expected winner
	}{
		{
			name: "no error",
			needError: false,
			inputData: context{100000, 100010},
			expected: winner{"Simple method is win!", 2, 0},
		},
		{
			name: "numbers must be positive",
			needError: true,
			inputData: context{-999999, 100000},
			expected: winner{"incorrect data!", 0, 0},
		},
		{
			name: "min must be lower than max and their length must be 6",
			needError: true,
			inputData: context{75, 50},
			expected: winner{"incorrect data!", 0, 0},
		},
		{
			name: "their length must be 6",
			needError: true,
			inputData: context{1234567, 12345678},
			expected: winner{"incorrect data!", 0, 0},
		},
	}

	for _, object := range testObjects{
		object := object
		t.Run(object.name, func(t *testing.T) {
			gotResult, err := GetLuckyTickets(&object.inputData)
			if object.needError != (err != nil){
				t.Errorf("for returned error: want <nil> but got %v", err)
			}
			if reflect.DeepEqual(gotResult, object.expected) == false{
				fmt.Println("expected struct: ", object.expected, "\nbut got:         ", gotResult)
			}
		})
	}



}