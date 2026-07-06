package mock

type MockParser struct {
}

func NewMockParser() *MockParser {
	return &MockParser{}
}

func (d *MockParser) ExtractText(path string) (string, error) {
	return "text" + path, nil
}
