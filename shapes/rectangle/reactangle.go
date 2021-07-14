package rectangle

type Rectangle struct {
	Length  float64 `json:"length"`
	Breadth float64 `json:"breadth"`
	Color   string  `json:"color"`
	Area    float64 `json:"area"`
}

func (rectangle *Rectangle) CalculateArea() {
	rectangle.Area = rectangle.Length * rectangle.Breadth
}
