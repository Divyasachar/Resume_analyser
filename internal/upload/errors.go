package upload

import (
	"errors"
)

var (
	ErrFileRequired      = errors.New("resume file is required")
	ErrInvalidFileType   = errors.New("only pdf, docx and txt files are allowed")
	ErrFileTooLarge      = errors.New("file size exceeds 5 MB")
)
