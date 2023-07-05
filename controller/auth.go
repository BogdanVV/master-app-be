package controller

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/bogdanvv/master-app-be/config/constants"
	"github.com/bogdanvv/master-app-be/models"
	"github.com/bogdanvv/master-app-be/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// TODO: add image upload like on user-update
func (c *Controller) Signup(ctx *gin.Context) {
	var input models.SignUpBody

	if err := ctx.ShouldBind(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	user, err := c.service.Signup(input.Name, input.Email, input.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// process the image
	if input.Image != nil {
		// TODO: move to UploadAvatar function in utils
		if input.Image.Size > int64(constants.PROFILE_IMAGE_MAX_SIZE) {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("acceptable file size is maximum %dMB", constants.PROFILE_IMAGE_MAX_SIZE/1024/1024)})
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

		input.Image.Filename = fmt.Sprintf("%s.%s", user.Id, fileType)
		dst := fmt.Sprintf("public/uploads/profile-images/%s", input.Image.Filename)
		if err := ctx.SaveUploadedFile(input.Image, dst); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to upload the image"})
			return
		}

		protocol := "http"
		if !strings.Contains(ctx.Request.Proto, "HTTP") {
			protocol = "https"
		}
		imageURL := fmt.Sprintf("%s://%s/profile-image/%s", protocol, ctx.Request.Host, input.Image.Filename)
		user, err = c.service.UpdateUser(user.Id, models.UserUpdateBody{ImageURL: imageURL})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	accessToken, err := utils.GenerateAccessToken(user.Id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"data": gin.H{"user": user, "accessToken": accessToken}})
}

func (c *Controller) Login(ctx *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invlaid body"})
		return
	}

	user, err := c.service.Login(input.Email, input.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": user})
}

func (c *Controller) RefreshToken(ctx *gin.Context) {
	authHeader := ctx.Request.Header.Get(constants.AUTHORIZATION_HEADER)
	token := strings.Split(authHeader, " ")[1]
	if token == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
		return
	}

	newToken, err := c.service.RefreshAccessTokenToken(token)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	jwtToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("ACESS_TOKEN_SECRET")), nil
	})
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok || !jwtToken.Valid {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	userId, _ := claims.GetSubject()
	user, err := c.service.GetUserById(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{"data": gin.H{"accessToken": newToken, "user": user}})
}
