package upload

import (
	"path/filepath"
	"strings"
)

const MaxFileSize = 5 * 1024 * 1024

var allowedExtensions = map[string]bool{
	".pdf":  true,
	".docx": true,
	".txt":  true,
}

func ValidateFile(fileSize int64, filename string) error {

	if fileSize > MaxFileSize {
		return ErrFileTooLarge
	}

	extension := strings.ToLower(filepath.Ext(filename))

	if !allowedExtensions[extension] {
		return ErrInvalidFileType
	}

	return nil
}