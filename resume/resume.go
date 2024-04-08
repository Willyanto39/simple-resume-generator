package resume

import "strings"

type Date struct {
	Month string
	Year  string
}

type Contact struct {
	Name    string
	Email   string
	Phone   string
	Address string
}

func (contact Contact) Information() string {
	contacts := []string{}

	if contact.Email != "" {
		contacts = append(contacts, contact.Email)
	}

	if contact.Phone != "" {
		contacts = append(contacts, contact.Phone)
	}

	if contact.Address != "" {
		contacts = append(contacts, contact.Address)
	}

	return strings.Join(contacts, " | ")
}

type Experience struct {
	Company  string
	Role     string
	Location string
	Start    Date
	End      Date
	Details  []string
}

type Education struct {
	Name     string
	Degree   string
	Location string
	Start    Date
	End      Date
	Details  []string
}

type Project struct {
	Name        string
	Link        string
	Details     []string
	Skills_Used []string
}

func (project Project) SkillsUsedString() string {
	return strings.Join(project.Skills_Used, ", ")
}

type Resume struct {
	Contact     Contact
	Experiences []Experience
	Educations  []Education
	Skills      []string
	Projects    []Project
}

func (resume Resume) SkillsString() string {
	return strings.Join(resume.Skills, ", ")
}
