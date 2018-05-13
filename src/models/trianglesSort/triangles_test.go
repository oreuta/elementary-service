package trianglesSort

import "testing"

func TestTrianglesSquareSort(t *testing.T) {

	testCases := []struct {
		name                   string
		input                  []Triangle
		expSortedVerticesNames []string
		needError              bool
	}{
		{
			name: "error, 3 triangles, japaneese symbol in vertices names",
			input: []Triangle{{"々B水", 10, 20, 22.36, 0},
				{"CB水", 16, 15, 6, 0},
				{"BAC", 10, 9, 16, 0}},
			needError: true,
		},

		{
			name: "error, 3 triangles, names of vertices dont fit its values",
			input: []Triangle{{"ABC", 10.0, 20.0, 22.36, 0},
				{"CBA", 16.0, 15.0, 6, 0},
				{"BAC", 16, 15, 6, 0}},
			needError: true,
		},
		{
			name: "no error, 5 triangles,4 of them have same names and values",
			input: []Triangle{{"ABC", 10.0, 20.0, 22.36, 0},
				{"CBA", 16.0, 15.0, 6, 0},
				{"CBA", 16.0, 15.0, 6, 0},
				{"CBA", 16.0, 15.0, 6, 0},
				{"CBA", 16.0, 15.0, 6, 0}},
			expSortedVerticesNames: []string{"CBA", "CBA", "CBA", "CBA", "ABC"},
			needError:              false,
		},
		{
			name: "no error, 2 triangles",
			input: []Triangle{{"ABC", 10, 20, 22.36, 0},
				{"CBA", 16, 15, 6, 0}},
			expSortedVerticesNames: []string{"CBA", "ABC"},
			needError:              false,
		},
		{
			name: "no error, 3 triangles",
			input: []Triangle{{"ABC", 10, 20, 22.36, 0},
				{"CBA", 16, 15, 6, 0},
				{"BAC", 10, 9, 16, 0}},
			expSortedVerticesNames: []string{"BAC", "CBA", "ABC"},
			needError:              false,
		},
		{
			name: "error, 3 triangles and some vertices are zero",
			input: []Triangle{{"ABC", 10, 20, 22.36, 0},
				{"CBA", 16, 0, 6, 0},
				{"BAC", 10, 9, 0, 0}},
			needError: true,
		},
		{
			name: "error, 3 triangles, one side is negative",
			input: []Triangle{{"ABC", 10.0, -20.0, 22.36, 0},
				{"CBA", 16.0, 15.0, 6, 0},
				{"BAC", 10.0, 9, 22.36, 16}},
			needError: true,
		},
		{
			name: "error, 2 triangles, 2nd have wrong triangle sizes",
			input: []Triangle{{"ABC", 10.0, 20.0, 22.36, 0},
				{"BAC", 10.0, 9, 22.36, 0}},
			needError: true,
		},
		{
			name:      "error, empty slice",
			input:     []Triangle{},
			needError: true,
		},
		{
			name:      "error, single triangle",
			input:     []Triangle{{"ABC", 10.0, 20.0, 22.36, 0}},
			needError: true,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			actualSlice, err := TrianglesSquareSort(testCase.input)

			if !testCase.needError {
				for i := range actualSlice {
					if actualSlice[i].Vertices != testCase.expSortedVerticesNames[i] {
						t.Errorf("for returned actualSlice: want %v but got %v", testCase.expSortedVerticesNames[i], actualSlice[i].Vertices)
					}
				}
			} else if testCase.needError != (err != nil) {
				t.Errorf("for returned error: want <nil> but got \"%v\"", err)
			}

		})
	}
}
