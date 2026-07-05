package matcher

import (
	"strings"

	"github.com/divya/resume-analyser/internal/job"
	"github.com/divya/resume-analyser/internal/resume"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}
func (s *Service) Match(
	resumeData *resume.Resume,
	jobData *job.JobDescription,
) *MatchResult {
	jobSkills := toSet(jobData.RequiredSkills)
	var matched []string
	var extra []string
	for _, skill := range resumeData.Skills {
	if jobSkills[strings.ToLower(skill)] {
			matched = append(
				matched,
				skill,
			)
		} else {
			extra = append(
				extra,
				skill,
			)
		}
	}
	resumeSkills := toSet(resumeData.Skills)
	var missing []string
	for _, skill := range jobData.RequiredSkills {
		if !resumeSkills[strings.ToLower(skill)] {
			missing = append(
				missing,
				skill,
			)
		}
	}
	score := 0.0
	if len(jobData.RequiredSkills) > 0 {
		score = float64(len(matched)) /
			float64(len(jobData.RequiredSkills)) *
			100
	}
	return &MatchResult{
		MatchedSkills: matched,
		MissingSkills: missing,
		ExtraSkills: extra,
		SkillScore: score,
	}
}