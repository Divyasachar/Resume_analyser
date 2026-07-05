package analyze

import (
	"fmt"
	"net/http"

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

func (h *Handler) Analyze(c *gin.Context) {

	path := c.Query("path")
	fmt.Println("path:", path)
	if path == "" {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "path is required",
		})

		return
	}

	response, err := h.service.Analyze(path)

	if err != nil {

		c.JSON(http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			})

		return
	}

	c.JSON(http.StatusOK, response)
}
