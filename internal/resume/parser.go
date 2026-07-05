package resume

import "github.com/divya/resume-analyser/internal/skills.go"

type Parser struct{}

func NewParser() *Parser {

	return &Parser{}
}

func (p *Parser) Parse(text string) *Resume {

	return &Resume{

		Name: skills.ExtractName(text),
		Email: skills.ExtractEmail(text),
		Phone: skills.ExtractPhone(text),
		Skills: skills.ExtractSkills(text),
	}
}