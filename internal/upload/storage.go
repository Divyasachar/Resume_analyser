package upload

import "mime/multipart"

type Storage interface {
	Save(file *multipart.FileHeader) (*UploadResponse, error)
}