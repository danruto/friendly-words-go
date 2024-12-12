package main

import (
	"fmt"

	fw "github.com/danruto/friendly-words-go/generator"
)

func main() {
	s := fw.GenerateUniqueWordListID(5)

	// Generates something like `chime-worried-potpourri-phalanx-assembly`
	fmt.Println(s)
}
