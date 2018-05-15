package squarert

import (
	"errors"
	"math"
)

// SquareRoot returns a slice of square root of its argument
func SquareRoot(numbers []int) (squareRoots []float64, err error) {
	squareRoots = make([]float64, len(numbers))
	for i, number := range numbers {
		if number < 0 {
			return []float64{}, errors.New("number must be positive")
		}
		squareRoots[i] = math.Sqrt(float64(number))
	}
	return
}

