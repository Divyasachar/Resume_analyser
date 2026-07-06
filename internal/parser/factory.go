package parser

import (
	"path/filepath"
	"strings"
)

func NewParser(path string) (Parser, error) {

	extension := strings.ToLower(
		filepath.Ext(path),
	)
	return Get(extension)
}