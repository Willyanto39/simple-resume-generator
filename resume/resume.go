package resume

import (
	"strings"
)

type Date struct {
	Month string
	Year  string
}

func (date Date) String() string {
	dateFields := []string{}

	if date.Month != "" {
		dateFields = append(dateFields, date.Month)
	}

	if date.Year != "" {
		dateFields = append(dateFields, date.Year)
	}

	return strings.Join(dateFields, " ")
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
	Company       string
	Role          string
	Location      string
	Until_Present bool
	Start         Date
	End           Date
	Details       []string
}

func (experience Experience) DateString() string {
	if experience.Until_Present {
		return experience.Start.String() + " - Present"
	}

	return experience.Start.String() + " - " + experience.End.String()
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
