package palindrome

import (
	"errors"
	"unicode"
)

// HasPalindrome returns a slice of all palindromes of its argument
func HasPalindrome(stringToExplore string) (palindromeSubstrings []string, err error) {
	if len([]rune(stringToExplore)) < 2 {
		return []string{}, errors.New("need to enter at least 2 symbols")
	}
	palindromeSubstrings = make([]string, 0)
	length := len([]rune(stringToExplore))

	for g, y := range stringToExplore {

		for substringLength := 1; substringLength <= length-g; substringLength++ {

			substring := stringToExplore[g : g+substringLength]
			if isPalindrome(substring) && len(substring) > 1 && unicode.IsLetter(y) {
				palindromeSubstrings = append(palindromeSubstrings, substring)
			}
		}
	}

	return
}
func isPalindrome(stringToTest string) bool {
	var letters []rune
	for _, r := range stringToTest {
		if unicode.IsLetter(r) {
			letters = append(letters, unicode.ToLower(r))
		}
	}
	for i := range letters {
		if letters[i] != letters[len(letters)-i-1] {
			return false
		}
	}
	return true

}
