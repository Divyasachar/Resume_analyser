package upload

import "mime/multipart"

type MockStorage struct {
	LastFile *multipart.FileHeader

	SaveCalled bool

	Path      string
	SaveCount int

	Err error
}

func (m *MockStorage) Save(file *multipart.FileHeader) (*UploadResponse, error) {

	m.SaveCalled = true
	m.LastFile = file
	m.SaveCount++

	return &UploadResponse{Path: m.Path}, m.Err
}
