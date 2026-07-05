package job

import "github.com/divya/resume-analyser/internal/skills.go"

type Parser struct{}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) Parse(text string) *JobDescription {

	requiredText, preferredText := skills.SplitSections(text)

	return &JobDescription{

		Title: skills.ExtractTitle(text),

		Description: text,

		RequiredSkills: skills.ExtractSkills(requiredText),

		PreferredSkills: skills.ExtractSkills(preferredText),
	}
}
