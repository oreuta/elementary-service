package fibo

import "testing"

func TestFibonacci(t *testing.T) {

	testCases := []struct {
		name      string
		input     []int
		fiboNumbers  []int
		needError bool
	}{
		{
			name:      		"no error, 2 positive numbers",
			input:     		[]int{1, 5},
			fiboNumbers:  	[]int{0, 1, 1, 2, 3, 5},
			needError: 		false,
		},
		{
			name:      "error, positive numbers + 1 negative",
			input:     []int{1, -1},
			needError: true,
		},
		{
			name:      "error, no numbers",
			input:     []int{},
			needError: true,
		},
		{
			name:      "error, to much numbers",
			input:     []int{1,2,3,4},
			needError: true,
		},
		{
			name:      "error, first number greater than second",
			input:     []int{5,1},
			needError: true,
		},
		{
			name:      "error, lenght of numbers < 1",
			input:     []int{0},
			needError: true,
		},
		{
			name:      "no error, lenght specified",
			input:     []int{8},
			fiboNumbers:  	[]int{14930352, 24157817, 39088169, 63245986},
			needError: false,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			actFiboNumbs, err := fibo.Fibo(testCase.input)
			if testCase.needError != (err != nil) { //XOR
				t.Errorf("for returned error: want <nil> but got %v", err)
			}
			for i := range testCase.fiboNumbers {
				if testCase.fiboNumbers[i] != actFiboNumbs[i] {
					t.Errorf("want %f but got %f", testCase.fiboNumbers[i], actFiboNumbs[i])
				}
			}
		})
	}
}
