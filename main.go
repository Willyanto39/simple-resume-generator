package main

import (
	"fmt"

	"github.com/Willyanto39/simple-resume-generator/converter"
	"github.com/Willyanto39/simple-resume-generator/parser"
)

func main() {
	fmt.Println("go")
	resume := parser.Parse()
	converter.ConvertToHtml(resume)
	converter.ConvertToPdf()
}
