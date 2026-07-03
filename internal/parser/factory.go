package parser

import (
	"errors"
	"path/filepath"
	"strings"
)

func NewParser(path string) (Parser, error) {

	switch strings.ToLower(filepath.Ext(path)) {

	case ".pdf":
		return NewPDFParser(), nil

	case ".docx":
		return NewDOCXParser(), nil

	case ".txt":
		return NewTXTParser(), nil

	default:
		return nil, errors.New("unsupported document type")
	}
}