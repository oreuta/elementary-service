package trianglesSort

import (
	"math"
	"sort"
	"errors"
	"regexp"
)

const (
	ERROR_SIGNED     = "Numbers cant be signed, or float or 0 or string"
	ERROR_WRONG_NAME = "Names must contain only ABC vertices, and triangle have only 3 vertices"
	ERROR_INPUT_SIZE = "Size of the input slice cant be 0 or 1. Can sort at least 2 triangles"
)

//Triangle is a structure, that holds Vertices names, their values and their sqrt
type Triangle struct {
	Vertices string  `json:"vertices"`
	A        float64 `json:"a"`
	B        float64 `json:"b"`
	C        float64 `json:"c"`
	Sqrt     float64 `json:"square,-"`
}

// TrianglesSquareSort is a function that sorts triangles slice by single triangle area. Returns sorted slice and error
func TrianglesSquareSort(trianglesToSortSlice []Triangle) ([]Triangle, error) {
	err := validateTriangles(trianglesToSortSlice)
	if err != nil {
		return trianglesToSortSlice, err
	}

	sort.Slice(trianglesToSortSlice, func(i, j int) bool {
		return trianglesToSortSlice[i].Sqrt < trianglesToSortSlice[j].Sqrt
	})

	return trianglesToSortSlice, nil
}

func validateTriangles(t []Triangle) (err error) {
	if len(t) == 0 || len(t) == 1 {
		return errors.New(ERROR_INPUT_SIZE)
	}
	for i := range t {
		err = t[i].validateSingleTriangle()
		if err != nil {
			return err
		}
		err = t[i].calculateSquare()
		if err != nil {
			return err
		}
	}
	return
}

func (t *Triangle) validateSingleTriangle() (err error) {
	var regex = regexp.MustCompile(`^[A-C]{3}$`) //своя реализация в анмаршалинге
	if regex.MatchString(t.Vertices) != true {
		return errors.New(ERROR_WRONG_NAME)
	}
	if t.A <= 0 || t.B <= 0 || t.C <= 0 {
		return errors.New(ERROR_SIGNED)
	}
	if !t.validateNamesFitNumbers() {
		return errors.New("Names dont fit actual values")
	}
	return
}

func (t *Triangle) validateNamesFitNumbers() bool {
	r := []rune(t.Vertices)
	if string(r[2]) == "A" && t.A > t.B && t.A > t.C {
		if string(r[0]) == "B" && t.B < t.C {
			return true
		} else if string(r[0]) == "C" && t.C < t.B {
			return true
		} else {
			return false
		}
	}
	if string(r[2]) == "B" && t.B > t.A && t.A > t.C {
		if string(r[0]) == "C" && t.C < t.A {
			return true
		} else if string(r[0]) == "A" && t.A < t.C {
			return true
		} else {
			return false
		}
	}
	if string(r[2]) == "C" && t.C > t.B && t.C > t.A {
		if string(r[0]) == "B" && t.B < t.A {
			return true
		} else if string(r[0]) == "A" && t.A < t.B {
			return true
		} else {
			return false
		}
	}
	return false
}

func (t *Triangle) calculateSquare() (err error) {
	p := 0.5 * (t.A + t.B + t.C)
	t.Sqrt = math.Sqrt(p * (p - t.A) * (p - t.B) * (p - t.C))
	if math.IsNaN(t.Sqrt) || t.Sqrt == 0 {
		return errors.New("Numbers are wrong for a triangle. Can not calculate square")
	}
	return
}
