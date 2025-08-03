package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/abuwho/md2html/utils"
)

func main() {
	source := flag.String("source", "./README.md", "Source of the markdown file")
	destination := flag.String("dest", "./result.html", "Where the result should be saved")

	flag.Parse()

	fmt.Println(*source, *destination)
	content, err := utils.ReadSourceFile(*source)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	utils.WriteToDestinationFile(*destination, content)
}
