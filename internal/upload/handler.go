package upload

import (
	"net/http"

	"github.com/divya/resume-analyser/internal/parser"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Upload(c *gin.Context) {
	getError := func(err error) {
		switch err {
		case ErrInvalidFileType:
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		case ErrFileTooLarge:
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "internal server error",
			})
		}
	}
	file, err := c.FormFile("resume")
	if err != nil {
		getError(err)
		return
	}
	response, err := h.service.Save(file)
	if err != nil {
		getError(err)
		return
	}
	p, err := parser.NewParser(response.Path)

	if err != nil {

		c.JSON(http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			})

		return
	}

	text, err := p.ExtractText(response.Path)
	if err != nil {

		c.JSON(http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			})

		return
	}

	response.Message = text
	c.JSON(http.StatusOK, response)
}
