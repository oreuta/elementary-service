package palindrome

import (
	"errors"
	"bytes"
)




var FindSubPalindromes = findSubPalindromes

//TestMockHandler() is a Mock function for findSubPalindromes function
func TestMockHandler() {
	FindSubPalindromes = func( s string) ([]string, error) {
		return []string{s}, nil
	}
}

// findSubPalindromes returns a string of all subpalindromes of its argument.
func findSubPalindromes(stringToExplore string) (resultString []string, err error) {
	var s string
	var bs bytes.Buffer
	var palindromeSubstrings [][]rune
	err = validate(stringToExplore)
	if err != nil {
		return resultString, err
	}
	length := len([]rune(stringToExplore))
	runedStr := []rune(stringToExplore)
	for offset := range stringToExplore {
		for substringLength := 1; substringLength <= length-offset; substringLength++ {
			substring := runedStr[offset: offset+substringLength]
			if isPalindrome(substring) && len(substring) > 1 {
				palindromeSubstrings = append(palindromeSubstrings, substring)
			}
		}
	}
	for j := range palindromeSubstrings {
		for i := range palindromeSubstrings[j] {
			bs.WriteString(string(palindromeSubstrings[j][i]))
		}
		s = bs.String()
		resultString = append(resultString, s)
		bs.Reset()
		
	}
	if len(resultString) == 0 {
		return
	}
	return
}

func isPalindrome(stringToTest []rune) bool {
	length := len(stringToTest)
	halfLength := length / 2

	for i, j := 0, length-1; i <= halfLength; i, j = i+1, j-1 {

		if stringToTest[i] != stringToTest[j] {
			return false
		}
	}
	return true
}
func validate(stringToExplore string) (error) {
	if len([]rune(stringToExplore)) < 2 {
		return  errors.New("need to enter at least 2 symbols")
	}
	return nil
}


