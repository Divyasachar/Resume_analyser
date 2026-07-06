package parser

import "fmt"

var registry = map[string]Registration{}

type Registration struct {
	Extension string

	Name string

	Parser Parser
}

func Register(name, extension string, parser Parser) {

	registry[extension] = Registration{Extension: extension, Name: name, Parser: parser}
}

func Get(extension string) (Parser, error) {

	p, ok := registry[extension]

	if !ok {
		return nil, fmt.Errorf("no parser registered for %s", extension)
	}

	return p.Parser, nil
}
