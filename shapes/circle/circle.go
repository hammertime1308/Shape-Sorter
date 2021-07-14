package circle

import (
	"math"
)

type Circle struct {
	Radius float64 `json:"radius"`
	Color  string  `json:"color"`
	Area   float64 `json:"area"`
}

func (circle *Circle) CalculateArea() {
	circle.Area = math.Pi * circle.Radius * circle.Radius
}
