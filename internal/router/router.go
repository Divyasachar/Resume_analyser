package router

import (
	"net/http"

	"github.com/divya/resume-analyser/internal/analyze"
	"github.com/divya/resume-analyser/internal/config"
	"github.com/divya/resume-analyser/internal/upload"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()

	storage := upload.NewLocalStorage()
	service := upload.NewService(storage)
	uploader := upload.NewHandler(service)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"application": config.AppConfig.AppName,
			"version":     config.AppConfig.AppVersion,
			"status":      "Running",
		})
	})

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "UP",
		})
	})

	r.GET("/version", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"version": config.AppConfig.AppVersion,
		})
	})

	analyzeService := analyze.NewService()
	analyzeHandler := analyze.NewHandler(analyzeService)
	r.POST("/upload", uploader.Upload)
	r.POST("/analyze", analyzeHandler.Analyze)
	return r
}
