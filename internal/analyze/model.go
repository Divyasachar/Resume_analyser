package analyze

import "github.com/divya/resume-analyser/internal/resume"

type Response struct {
	Resume *resume.Resume `json:"resume"`
}

var (
	ATSScore int

	Suggestions []string

	MatchedSkills []string

	MissingSkills []string
)
