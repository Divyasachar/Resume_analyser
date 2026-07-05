package ats

type Result struct {
	Score      int      `json:"score"`
	Grade      string   `json:"grade"`
	Badge      string   `json:"badge"`
	Strengths  []string `json:"strengths"`
	Weaknesses []string `json:"weaknesses"`
}
type GradeRule struct {
	Minimum int
	Grade   string
}

var Grades []GradeRule = []GradeRule{
	{90, "A+"},
	{80, "A"},
	{70, "B"},
	{60, "C"},
	{50, "D"}}
