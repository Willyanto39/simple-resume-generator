package main

import (
	"fmt"

	"github.com/Willyanto39/simple-resume-generator/parser"
)

func main() {
	fmt.Println("go")
	resume := parser.Parse()
	fmt.Println(resume)
}
