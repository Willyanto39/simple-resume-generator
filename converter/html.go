package converter

import (
	"fmt"
	"html/template"
	"os"

	"github.com/Willyanto39/simple-resume-generator/resume"
)

func ConvertToHtml(resume resume.Resume) {
	tmpl := template.Must(template.ParseFiles("resume-template.html"))
	tmpl.Option("missingkey=error")

	file, err := os.Create("index.html")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = tmpl.Execute(file, resume)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
