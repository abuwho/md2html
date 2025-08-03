package main

import (
	"flag"
	"fmt"
)

func main() {
	source := flag.String("source", "./README.md", "Source of the markdown file")
	destination := flag.String("dest", "./result.html", "Where the result should be saved")

	flag.Parse()

	fmt.Println(*source, *destination)
}
