package cli

import (
	"fmt"

	"github.com/hammertime1308/shape-sorter/shapes"
)

func Solve(input string) {
	// convert to string logic(cli hence already string)
	object := shapes.Parse(input)
	object.Sort()
	out := shapes.Serialize(object)
	fmt.Println(out)
}
