package triangle

import "math"

// Kind of a triangle according to its sides
type Kind string

const (
	NaT Kind = "not a triangle"
	Equ Kind = "equilateral"
	Iso Kind = "isosceles"
	Sca Kind = "scalene"
)

// KindFromSides returns the kind of a triangle according to its sides
func KindFromSides(a, b, c float64) Kind {
	// All sides have to be numbers
	if math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) {
		return NaT
	}
	// Sort sides by value
	if b < a {
		a, b = b, a
	}
	if c < b {
		b, c = c, b
	}
	if b < a {
		a, b = b, a
	}
	switch {
	// All sides have to be strictly positive
	case a <= 0:
		return NaT
	// Respect triangle equality (a + b >= c)
	case a+b < c:
		return NaT
	// The largest side cant be Inf to avoid the case where x, y, z = (z, +Inf, +Inf)
	case math.IsInf(c, 1):
		return NaT
	case a == c:
		return Equ
	case a == b || b == c:
		return Iso
	}
	return Sca
}
