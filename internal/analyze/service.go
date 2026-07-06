package analyze

import (
	"github.com/divya/resume-analyser/internal/parser"
	"github.com/divya/resume-analyser/internal/resume"
)

type Service struct {
	resumeParser *resume.Parser
}

func NewService() *Service {
	return &Service{
		resumeParser: resume.NewParser(),
	}
}
func (s *Service) Analyze(path string) (*Response, error) {

	documentParser, err := parser.NewParser(path)

	if err != nil {
		return nil, err
	}

	text, err := documentParser.ExtractText(path)

	if err != nil {
		return nil, err
	}

	resumeData := s.resumeParser.Parse(text)

	return &Response{
		Resume: resumeData,
	}, nil
}
