package palindrome

import (
	"errors"
	"strings"
)

// SubPalindromes returns a string of all subpalindromes of its argument
func SubPalindromes(stringToExplore string) (resultString string, err error) {
	stringToExplore = strings.Replace(stringToExplore, " ", "", -1)
	var palindromeSubstrings [][]rune
	err=valid(stringToExplore)
	if err!= nil{
		return stringToExplore, err
	}
	length := len([]rune(stringToExplore))
	runedStr := []rune(stringToExplore)
	for g, _ := range stringToExplore {
		for substringLength := 1; substringLength <= length-g; substringLength++ {
			substring := runedStr[g : g+substringLength]
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
	return
}

func isPalindrome(stringToTest []rune) bool {
	length := len(stringToTest)
	halfLength := length / 2

	for i := 0; i <= halfLength; i++ {
		if stringToTest[i] != stringToTest[length-i-1] {
			return false
		}
	}
	return true
}
func valid (stringToExplore string) (err error){
	if len([]rune(stringToExplore)) < 2 {
		return errors.New("need to enter at least 2 symbols")
	}
	return nil
}

