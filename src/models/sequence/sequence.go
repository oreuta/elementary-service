package sequence

import (
	"errors"
	"fmt"
	"strings"
)

//ValidateTheInput validates the input data and ensures that it doesn't contain any separators, symbols etc
//that may endanger the correct work of the algorithm
func ValidateTheInput(inputLength, inputMinSquare int) (outputLength, outputMinSquare int, err error) {
	if inputLength < 0 || inputMinSquare < 0 {
		return 0, 0, errors.New("Negative numbers used as an input data")
	}
	if inputLength == 0 || inputMinSquare == 0 || inputLength == 0 && inputMinSquare == 0 {
		return 0, 0, errors.New("One or both values is/are 0")
	}
		_, err = GetSquares(outputLength, outputMinSquare)
	return
}

//GetSquares builds an infinite loop with the sequence of natural digits those square
// is less than the one specified in the input data
func GetSquares(outputLength, outputMinSquare int) (arr []int, err error) {
	var j, sq int
	for {
		sq = j * j
		if sq >= outputMinSquare {
			if len(arr) == outputLength {
				break
			}
			arr = append(arr, sq)
		}
		j++
	}
	if err != nil {
		return arr, errors.New("Program exited with error")
	}
	return
}

//PrintWithCommas represents the array of digits as a string with commas between the numbers
func PrintWithCommas(arr []int) (string) {
	return strings.Replace(strings.Trim(fmt.Sprintf("%v", arr), "[]"), " ", ", ", -1)
}
