package parser

import (
	"github.com/ledongthuc/pdf"
)

type PDFParser struct{}

func NewPDFParser() *PDFParser {
	return &PDFParser{}
}

func (p *PDFParser) ExtractText(path string) (string, error) {

	file, reader, err := pdf.Open(path)

	if err != nil {
		return "", err
	}

	defer file.Close()

	var text string

	total := reader.NumPage()

	for i := 1; i <= total; i++ {

		page := reader.Page(i)

		if page.V.IsNull() {
			continue
		}

		content, err := page.GetPlainText(nil)

		if err != nil {
			return "", err
		}

		text += content
	}

	return text, nil
}