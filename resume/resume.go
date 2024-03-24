package resume

type Date struct {
	Month string
	Year  string
}

type Contact struct {
	Name  string
	Email string
	Phone string
}

type Experience struct {
	Company string
	Role    string
	Start   Date
	End     Date
}

type Education struct {
	Name        string
	Start       Date
	End         Date
	Description string
}

type Project struct {
	Name        string
	Description string
	Skills_Used []string
}

type Resume struct {
	Contact     Contact
	Experiences []Experience
	Educations  []Education
	Skills      []string
	Projects    []Project
}
