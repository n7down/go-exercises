package triangle

import "errors"

type Triangle struct {
	Side1 int
	Side2 int
	Side3 int
}

func NewTriangle(side1, side2, side3 int) *Triangle {
	return &Triangle{
		Side1: side1,
		Side2: side2,
		Side3: side3,
	}
}

func (t Triangle) validate() bool {
	if t.Side1 < 1 || t.Side2 < 1 || t.Side3 < 1 {
		return false
	}

	if t.Side1+t.Side2 > t.Side3 {
		return false
	}

	if t.Side2+t.Side3 > t.Side1 {
		return false
	}

	if t.Side1+t.Side3 > t.Side2 {
		return false
	}
	return true
}

func (t Triangle) TriangleType() (string, error) {
	valid := t.validate()
	if !valid {
		return "", errors.New("not a triangle")
	}

	if t.Side1 == t.Side2 && t.Side2 == t.Side3 && t.Side1 == t.Side3 {
		return "equalateral", nil
	} else if t.Side1 == t.Side2 || t.Side2 == t.Side3 || t.Side1 == t.Side3 {
		return "isosceles", nil
	} else {
		return "scalar", nil
	}
}
