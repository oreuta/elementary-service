package sequence

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

//Function Checkthedata validates the input data and ensures that it doesn;t contain any separators, symbols etc
func Checkthedata(lenght string, minSq string) (lenght1 int, minSq1 int, err error) {
	r := regexp.MustCompile("^[0-9]+$")
	if r.MatchString(minSq) != true || r.MatchString(lenght) != true ||
		r.MatchString(minSq) != true && r.MatchString(lenght) != true {
		fmt.Println("Error. You entered a letter or a number with a separator.Program exited")
		return 0, 0, errors.New("Input data contain symbols")
	}

	lenght1, _ = strconv.Atoi(lenght)
	minSq1, _ = strconv.Atoi(minSq)
	if lenght1 < 0 || minSq1 < 0 {
		fmt.Println("You entered a negative number. Program excited")
		return lenght1, minSq1, errors.New("Negative numbers used as an input data")
	}
	if lenght1 == 0 || minSq1 == 0 || lenght1 == 0 && minSq1 == 0 {
		fmt.Println("Either of two values equals to 0")
		return lenght1, minSq1, errors.New("One or both values is/are 0")
	} else {
		GetSquares(lenght1, minSq1)
	}
	return
}

/*func main() {
	ll1, ms1, err := Checkthedata("4", "2")
	arr, err := GetSquares(ll1, ms1)
	str := Printwithcommas(arr)
	fmt.Println("Numerical sequence separated by commas:  ", str)
	fmt.Println(err)
}*/

//Function in an infinite loop builds the sequence of natural digits those square is less than the one specified in the input data
func GetSquares(lenght1, minSq1 int) (arr []int, err error) {
	var sq int
	j := 0
	arr = []int{}
	for {
		sq = j * j
		if sq > minSq1 {
			if len(arr) == lenght1 {
				break
			}
			arr = append(arr, sq)
		}
		j++
	}
	if err != nil {
		fmt.Println("Program exited with error", err)
	}
	return
}

//Function Printwithcommas represents the array of digits as a string with commas between the numbers
func Printwithcommas(arr []int) (string) {
	if len(arr) == 0 {
		return " "
	}
	estimate := len(arr) * 4
	b := make([]byte, 0, estimate)
	for _, n := range arr {
		b = strconv.AppendInt(b, int64(n), 10)
		b = append(b, ',')
	}
	b = b[:len(b)-1]
	return string(b)
}
