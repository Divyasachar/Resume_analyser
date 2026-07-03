package upload

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler() *Handler {
	return &Handler{
		service: NewService(),
	}
}

func (h *Handler) Upload(c *gin.Context) {

	file, err := c.FormFile("resume")

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "resume file is required",
		})

		return
	}

	response, err := h.service.Save(file)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, response)
}