package upload

type UploadResponse struct {
	ID       string `json:"id"`
	FileName string `json:"file_name"`
	Size     int64  `json:"size"`
	Message  string `json:"message"`
	Path     string `json:"path"`
}
