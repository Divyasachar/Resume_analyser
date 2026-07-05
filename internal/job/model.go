package job

type JobDescription struct {
	Title string `json:"title"`

	Description string `json:"description"`

	RequiredSkills []string `json:"required_skills"`

	PreferredSkills []string `json:"preferred_skills"`
}