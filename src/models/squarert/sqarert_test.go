package squarert

import "testing"

func TestSquareRoot(t *testing.T) {

	testCases := []struct {
		name      string
		input     []int
		expRoots  []float64
		needError bool
	}{
		{
			name:      "no error, 3 positive numbers",
			input:     []int{1, 2, 3},
			expRoots:  []float64{1, 1.4142135623730951, 1.7320508075688772},
			needError: false,
		},
		{
			name:      "error, 3 positive numbers + 1 negative",
			input:     []int{1, 2, 3, -1},
			needError: true,
		},
		{},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T){
			actRoots, err := SquareRoot(testCase.input)
			if testCase.needError != (err != nil) { //XOR
				t.Errorf("for returned error: want <nil> but got %v", err)
			}
			for i := range testCase.expRoots {
				if testCase.expRoots[i] != actRoots[i] {
					t.Errorf("want %f but got %f", testCase.expRoots[i], actRoots[i])
				}
			}
		})
	}
}
