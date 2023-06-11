package controller

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/bogdanvv/master-app-be/models"
	"github.com/gin-gonic/gin"
)

func (c *Controller) UpdateUser(ctx *gin.Context) {
	var input models.UserUpdateBody
	userId := ctx.Param("id")
	if err := ctx.ShouldBind(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "could not parse data"})
		return
	}
	isBodyEmpty := input.Email == "" && input.Image == nil && input.Name == "" && input.Password == ""
	if isBodyEmpty {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "enpty update body"})
		return
	}

	maxFileSize := 5 << 20 // 5MB
	if input.Image.Size > int64(maxFileSize) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("acceptable file size is maximum %dMB", maxFileSize/1024/1024)})
	}
	contentTypeImageHeader := input.Image.Header.Get("Content-Type") // image/jpeg, image/jpg, image/png/, ...
	contentDispositionImageHeader := input.Image.Header.Get("Content-Disposition")
	contentDispositionImageHeaderChunks := strings.Split(contentDispositionImageHeader, ";")
	var fileType string // jpeg, jpg, png, ...
	for _, c := range contentDispositionImageHeaderChunks {
		if strings.Contains(c, "filename") {
			_, after, _ := strings.Cut(c, `filename="`)
			fileName := strings.TrimRight(after, `"`)
			fileTypeChunks := strings.Split(fileName, ".")
			fileType = fileTypeChunks[len(fileTypeChunks)-1]
			break
		}
	}
	acceptableFileTypes := []string{"jpeg", "jpg", "png", "image/jpeg", "image/jpg", "image/png"}
	var isTypeOK bool
	for _, e := range acceptableFileTypes {
		if fileType == e || contentTypeImageHeader == e {
			isTypeOK = true
			break
		}
	}
	if !isTypeOK {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("invalid file type, only %s types are acceptable", strings.Join(acceptableFileTypes, ", "))})
		return
	}

	input.Image.Filename = fmt.Sprintf("%s.%s", userId, fileType)
	dst := fmt.Sprintf("public/uploads/profile-images/%s", input.Image.Filename)
	if err := ctx.SaveUploadedFile(input.Image, dst); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to upload the image"})
		return
	}

	protocol := "http"
	if !strings.Contains(ctx.Request.Proto, "HTTP") {
		protocol = "https"
	}
	imageUrl := fmt.Sprintf("%s://%s/profile-image/%s", protocol, ctx.Request.Host, input.Image.Filename)
	input.ImageURL = imageUrl

	user, err := c.service.UpdateUser(userId, input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": user})
}

func (c *Controller) GetUserById(ctx *gin.Context) {
	userId := ctx.Param("id")
	user, err := c.service.GetUserById(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}
