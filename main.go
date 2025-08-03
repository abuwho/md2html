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

	// Verify file type
	_, sourceFileTypeVerifiedErr := utils.VerifySourceFileType(*source)
	if sourceFileTypeVerifiedErr != nil {
		fmt.Println("Error: ", sourceFileTypeVerifiedErr)
		os.Exit(1)
	}

	fileContent, readFileErr := utils.ReadSourceFile(*source)
	if readFileErr != nil {
		fmt.Println("Error: ", readFileErr)
		os.Exit(1)
	}

	utils.WriteToDestinationFile(*destination, fileContent)
}
