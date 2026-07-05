package matcher

type MatchResult struct {
	MatchedSkills []string `json:"matched_skills"`

	MissingSkills []string `json:"missing_skills"`

	ExtraSkills []string `json:"extra_skills"`

	SkillScore float64 `json:"skill_score"`
}