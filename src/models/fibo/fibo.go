package fibo

import (
	"errors"
	"math"
)

func _fib(numIndex int) int{

		var result int
		for i, first, second := 0, 0, 1; i <= numIndex; i, first, second = i+1, first+second, first {
		if i == numIndex {
		result = first
	}
	}
		return result
	}


func lenghtOfNumber(numberToProceed int) (lenghtOfNumber int) {
	if numberToProceed == 0 {
		return 1
	} else {
		return int(math.Ceil(math.Log10(math.Abs(float64(numberToProceed)) + 0.5)))
	}
}

// Fibo returns a slice of Fibonacci numbers depending to argument
func Fibo(numbers []int) (fiboNumbers []int, err error) {

	workMode :=0			// 1 - requested lenght, 2- requested range

	for i:= 0; i< len(numbers); i++ {
		if numbers[i] < 0 {
			return []int{}, errors.New("using of negative numbers is not allowed")
		}
	}

	if len(numbers) == 1 {
		if numbers[0] < 1 {
			return []int{}, errors.New("lenght of result(s) number can't be less of 1 digit")
		}
		workMode = 1       	// result numbers must have requested lenght
	} else if len(numbers) == 2 {
		if numbers[0] >= numbers[1] {
			return []int{}, errors.New("start of range can't be greater or equal the finish one")
		}
		workMode = 2		// result numbers must be in requested range
	} else {
		return []int{}, errors.New("input data not recognized")
	}


	rangeParsed := false
	fiboNumbers = make([]int, 0)
	fiboNumberTemp := 0
	tempNumber := 0

	if workMode == 2 {

		for !rangeParsed {
			fiboNumberTemp = _fib(tempNumber)
			if numbers[0] <= fiboNumberTemp && fiboNumberTemp <= numbers[1] {
				fiboNumbers = append(fiboNumbers, fiboNumberTemp)
			}
			if fiboNumberTemp > numbers[1] {
				rangeParsed = true
			}
			tempNumber++
		}
		return fiboNumbers, nil // ------->  return sequence of fibnumbers in requested range

	} else if workMode == 1 {

		for !rangeParsed {
			fiboNumberTemp = _fib(tempNumber)
			if numbers[0] == lenghtOfNumber(fiboNumberTemp) {
				fiboNumbers = append(fiboNumbers, fiboNumberTemp)
			}
			if lenghtOfNumber(fiboNumberTemp) > numbers[0] {
				rangeParsed = true
			}
			tempNumber++
		}
		return fiboNumbers, nil // ------->  return sequence of fibnumbers of requested lenght
	}
	return
}
