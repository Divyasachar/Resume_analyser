package upload

import "mime/multipart"

type Service struct {
	storage Storage
}

func NewService(storage Storage) *Service {
	return &Service{
		storage: storage,
	}
}

func (s *Service) Save(file *multipart.FileHeader) (*UploadResponse, error) {

	err := ValidateFile(file.Size, file.Filename)
	if err != nil {
		return nil, err
	}

	return s.storage.Save(file)
}