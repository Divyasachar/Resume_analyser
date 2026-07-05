package ats

import "github.com/divya/resume-analyser/internal/matcher"

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Calculate(
	match *matcher.MatchResult,
) *Result {

	score := int(match.SkillScore)

	var strengths []string
	var weaknesses []string

	if len(match.MatchedSkills) > 0 {
		strengths = append(
			strengths,
			"Relevant technical skills found",
		)
	}

	for _, skill := range match.MissingSkills {

		weaknesses = append(
			weaknesses,
			"Missing "+skill,
		)
	}
	return &Result{
		Badge:      badge(score),
		Score:      score,
		Grade:      grade(score),
		Strengths:  strengths,
		Weaknesses: weaknesses,
	}
}
func grade(score int) string {
	for _, grade := range Grades {
		if score >= grade.Minimum {
			return grade.Grade
		}
	}
	return "F"
}
func badge(score int) string {

	if score >= 95 {
		return "Perfect Match"
	}

	return ""
}
