package txt

import (
	"os"

	"github.com/divya/resume-analyser/internal/parser"
)

type TXTParser struct{}

func NewTXTParser() *TXTParser {
	return &TXTParser{}
}

func (t *TXTParser) ExtractText(path string) (string, error) {

	data, err := os.ReadFile(path)

	if err != nil {
		return "", err
	}

	return string(data), nil
}
func init() {
	parser.Register("txt", ".txt", &TXTParser{})
}
