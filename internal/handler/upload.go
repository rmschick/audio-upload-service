package upload

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"personal-dev/internal/service"
)

// CreateUploadHandler handles file upload requests.
func CreateUploadHandler(c *gin.Context, fs service.FileService) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing file"})

		return
	}

	req := service.CreateFileRequest{
		File:       file,
		DeviceID:   c.PostForm("device_id"),
		Location:   c.PostForm("location"),
		UploadedAt: time.Now(),
	}

	f, err := fs.CreateFile(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusOK, f)
}

// GetUploadHandler retrieves a file by its ID.
func GetUploadHandler(c *gin.Context, fs service.FileService) {
	fileID := c.Param("id")

	f, err := fs.GetFile(c.Request.Context(), fileID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusOK, f)
}
