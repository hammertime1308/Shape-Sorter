package shapes

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"

	"github.com/hammertime1308/shape-sorter/shapes/circle"
	"github.com/hammertime1308/shape-sorter/shapes/rectangle"
)

type Shapes struct {
	Rectangles []*rectangle.Rectangle `json:"rectangles"`
	Circles    []*circle.Circle       `json:"circles"`
}

func Parse(input string) *Shapes {
	object := Shapes{}
	err := json.Unmarshal([]byte(input), &object)
	if err != nil {
		fmt.Println("Error = ", err)
		os.Exit(2)
	}
	if len(object.Rectangles) != 0 {
		for _, rectangle := range object.Rectangles {
			rectangle.CalculateArea()
		}
	}
	if len(object.Circles) != 0 {
		for _, circle := range object.Circles {
			circle.CalculateArea()
		}
	}
	return &object
}

func Serialize(object *Shapes) string {
	out, err := json.MarshalIndent(object, "", "\t")
	if err != nil {
		fmt.Println("Error = ", err.Error())
		os.Exit(2)
	}
	return string(out)
}

func (object *Shapes) Sort() {
	sort.Slice(object.Rectangles, func(i, j int) bool {
		return object.Rectangles[i].Area < object.Rectangles[j].Area
	})
	sort.Slice(object.Circles, func(i, j int) bool {
		return object.Circles[i].Area < object.Circles[j].Area
	})
}
