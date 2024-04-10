package main

import (
	"github.com/Willyanto39/simple-resume-generator/converter"
	"github.com/Willyanto39/simple-resume-generator/parser"
)

func main() {
	resume := parser.Parse()
	converter.ConvertToHtml(resume)
	converter.ConvertToPdf()
}
