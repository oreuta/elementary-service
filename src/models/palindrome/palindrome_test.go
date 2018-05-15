package palindrome

import "testing"

func TestFindSubPalindromes(t *testing.T) {

	testCases := []struct {
		name           string
		input          string
		expPalindromes string
		needError      bool
	}{
		{
			name:           "no error, testing 'taat'",
			input:          "taat",
			expPalindromes: "taat aa ",
			needError:      false,
		},
		{
			name:      "error, length of entered string is 0",
			input:     " ",
			needError: true,
		},
		{name: "error, length of entered string is 1",
			input: "j",
			needError: true,},

		{name: "error, test one Chinese symbol",
			input: "世",
			needError: true,
		},

		{name: "error, test two Chinese symbols",
			input: "世界",
			needError: false,
		},
		{
			name:           "no error, testing 'kayak'",
			input:          "kayak",
			expPalindromes:"kayak aya ",
			needError:      false,
		},
	}
	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			actPalindromes, err := FindSubPalindromes(testCase.input)
			if testCase.needError != (err != nil) {
				t.Errorf("for returned error: want <nil> but got %v", err)
			}
			for i := range testCase.expPalindromes {
				if testCase.expPalindromes[i] != actPalindromes[i] {
					t.Errorf("want %s but got %s", testCase.expPalindromes[i], actPalindromes[i])
				}
			}
		})
	}
}

