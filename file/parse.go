package file

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/hammertime1308/shape-sorter/shapes"
)

func Solve(path string) {
	// read file in the path and do operations
	fileData, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("Error = ", err.Error())
		os.Exit(2)
	}

	object := shapes.Parse(string(fileData))
	object.Sort()
	out := shapes.Serialize(object)
	fmt.Println(out)
}
