package parser

import (
	"github.com/nguyenthenguyen/docx"
)

type DOCXParser struct{}

func NewDOCXParser() *DOCXParser {
	return &DOCXParser{}
}

func (d *DOCXParser) ExtractText(path string) (string, error) {

	doc, err := docx.ReadDocxFile(path)

	if err != nil {
		return "", err
	}

	defer doc.Close()

	text := doc.Editable().GetContent()

	return text, nil
}