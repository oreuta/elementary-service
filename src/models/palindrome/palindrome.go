package palindrome

import (
	"errors"
	"strings"
)

// FindSubPalindromes returns a string of all subpalindromes of its argument.
func FindSubPalindromes(stringToExplore string) (resultString string, err error) {
	stringToExplore = strings.Replace(stringToExplore, " ", "", -1)
	var palindromeSubstrings [][]rune
	err = validate(stringToExplore)
	if err != nil {
		return stringToExplore, err
	}
	length := len([]rune(stringToExplore))
	runedStr := []rune(stringToExplore)
	for offset, _ := range stringToExplore {
		for substringLength := 1; substringLength <= length-offset; substringLength++ {
			substring := runedStr[offset: offset+substringLength]
			if isPalindrome(substring) && len(substring) > 1 {
				palindromeSubstrings = append(palindromeSubstrings, substring)
			}
		}
	}
	for j := range palindromeSubstrings {
		for i := range palindromeSubstrings[j] {
			resultString += string(palindromeSubstrings[j][i])
		}
		resultString += " "
	}

	if len(resultString) == 0 {
		return
	}
	resultRune := []rune (resultString)
	resultString = string(resultRune[0:len(resultRune)-1])
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
func validate(stringToExplore string) (err error) {
	if len([]rune(stringToExplore)) < 2 {
		return errors.New("need to enter at least 2 symbols")
	}
	return nil
}
