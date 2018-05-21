package trianglesSort

import (
	"math"
	"sort"
	"errors"
	"regexp"
)

const (
	errBadFormat   = "Numbers can't be signed, or float or 0 or string"
	errorWrongName = "Names must contain only ABC vertices, and triangle have only 3 vertices"
	errorInputSize = "Size of the input slice cant be 0 or 1. Can sort at least 2 triangles"
)

//Triangle is a structure, that holds Vertices names, their values and their sqrt
type Triangle struct {
	Vertices string  `json:"vertices"`
	A        float64 `json:"a"`
	B        float64 `json:"b"`
	C        float64 `json:"c"`
	Sqrt     float64 `json:"square,-"`
}

var regex = regexp.MustCompile(`^[A-C]{3}$`)


//TrianglesSquareSort is a function that sorts triangles slice by single triangle area. Returns sorted slice and error
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
		return errors.New(errorInputSize)
	}
	for i := range t {
		err = t[i].validateSingleTriangle()
		if err != nil {
			return
		}
		err = t[i].calculateSquare()
		if err != nil {
			return
		}
	}
	return
}

func (t *Triangle) validateSingleTriangle() (err error) {

	if !regex.MatchString(t.Vertices) {
		return errors.New(errorWrongName)
	}
	if t.A <= 0 || t.B <= 0 || t.C <= 0 {
		return errors.New(errBadFormat)
	}
	if !t.validateNamesFitNumbers() {
		return errors.New("Names don't fit actual values")
	}
	return
}

func (t *Triangle) validateNamesFitNumbers() (flag bool) {
	r := []rune(t.Vertices)
	if t.validateNamedVerticeA(r) {
		return true
	} else if t.validateNamedVerticeB(r) {
		return true
	} else if t.validateNamedVerticeC(r) {
		return true
	}
	return
}

func (t *Triangle) validateNamedVerticeA(r []rune) (flag bool) {
	if string(r[2]) == "A" && t.A > t.B && t.A > t.C {
		if string(r[0]) == "B" && t.B < t.C {
			return true
		} else if string(r[0]) == "C" && t.C < t.B {
			return true
		} else {
			return
		}
	}
	return
}

func (t *Triangle) validateNamedVerticeB(r []rune) (flag bool) {
	if string(r[2]) == "B" && t.B > t.A && t.A > t.C {
		if string(r[0]) == "C" && t.C < t.A {
			return true
		} else if string(r[0]) == "A" && t.A < t.C {
			return true
		} else {
			return
		}
	}
	return
}

func (t *Triangle) validateNamedVerticeC(r []rune) (flag bool) {
	if string(r[2]) == "C" && t.C > t.B && t.C > t.A {
		if string(r[0]) == "B" && t.B < t.A {
			return true
		} else if string(r[0]) == "A" && t.A < t.B {
			return true
		} else {
			return
		}
	}
	return
}

func (t *Triangle) calculateSquare() (err error) {
	p := 0.5 * (t.A + t.B + t.C)
	t.Sqrt = math.Sqrt(p * (p - t.A) * (p - t.B) * (p - t.C))
	if math.IsNaN(t.Sqrt) || t.Sqrt == 0 {
		return errors.New("Numbers are wrong for a triangle. Can not calculate square")
	}
	return
}
