package sequence

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//Function Validatetheinput validates the input data and ensures that it doesn't contain any separators, symbols etc
//that may endanger the correct work of the algorithm

func Validatetheinput(lenghtofthesequence, minsquare string,r *regexp.Regexp) (lenghtofthesequence1, minsquare1 int, err error) {
	if !r.MatchString(minsquare) || !r.MatchString(lenghtofthesequence)  ||
		!r.MatchString(minsquare) && !r.MatchString(lenghtofthesequence)  {
		return 0, 0, errors.New("Input data contain symbols or a separator")
	}
	lenghtofthesequence1, _ = strconv.Atoi(lenghtofthesequence)
	minsquare1, _ = strconv.Atoi(minsquare)
	if lenghtofthesequence1 < 0 || minsquare1 < 0 {
		return 0, 0, errors.New("Negative numbers used as an input data")
	}
	if lenghtofthesequence1 == 0 || minsquare1 == 0 || lenghtofthesequence1 == 0 && minsquare1 == 0 {
		return 0, 0, errors.New("One or both values is/are 0")
	} else {
		GetSquares(lenghtofthesequence1, minsquare1)
	}
	return
}

//getSquares builds an infinite loop with the sequence of natural digits those square
// is less than the one specified in the input data
func GetSquares(lenghtofthesequence1, minsquare1 int) (arr []int, err error) {
	var j, sq int
	for {
		sq = j * j
		if sq > minsquare1 {
			if len(arr) == lenghtofthesequence1 {
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

//Function PrintWithCommas represents the array of digits as a string with commas between the numbers
func PrintWithCommas(arr []int) (string) {
	return strings.Replace(strings.Trim(fmt.Sprintf("%v", arr), "[]"), " ", ", ", -1)
}
