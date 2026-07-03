package upload

import (
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Save(file *multipart.FileHeader) (*UploadResponse, error) {

	id := uuid.New().String()

	ext := filepath.Ext(file.Filename)

	newName := fmt.Sprintf("%s%s", id, ext)

	savePath := filepath.Join("uploads", newName)

	if err := os.MkdirAll("uploads", os.ModePerm); err != nil {
		return nil, err
	}

	if err := saveUploadedFile(file, savePath); err != nil {
		return nil, err
	}

	return &UploadResponse{
		ID:       id,
		FileName: newName,
		Size:     file.Size,
		Message:  "Upload successful",
	}, nil
}