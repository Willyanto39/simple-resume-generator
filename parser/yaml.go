package parser

import (
	"fmt"
	"os"

	"github.com/Willyanto39/simple-resume-generator/resume"
	"gopkg.in/yaml.v3"
)

func Parse() resume.Resume {
	file, err := os.ReadFile("resume.yaml")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	resume_data := resume.Resume{}
	err = yaml.Unmarshal(file, &resume_data)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return resume_data
}
