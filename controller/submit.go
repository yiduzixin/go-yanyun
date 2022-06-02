package controller

import (
	"dropzone-go/libs"
	"dropzone-go/services"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"net/http"
)

func ShowIndex(c *gin.Context) {
	masterUuid := uuid.NewV4()
	logoUuid := uuid.NewV4()
	c.HTML(http.StatusOK, "index.html", gin.H{
		"logo_uuid":   logoUuid.String(),
		"master_uuid": masterUuid.String(),
	})
}

func GetFormData(c *gin.Context) error {
	uploadType := c.PostForm("upload_type")
	accessKey := c.PostForm("access_key")
	secretKey := c.PostForm("secret_key")
	logoUuid := c.PostForm("logo_uuid")
	masterUuid := c.PostForm("master_uuid")
	err := libs.ValidateAccess(accessKey, secretKey)
	if err != nil {
		return libs.AccessForbidden()
	}
	dateString := libs.GetDateFormat()

	fileneme := services.Save2cloud(uploadType, dateString, masterUuid, logoUuid)

	r := services.PrettyData(uploadType, fileneme, dateString)
	c.JSON(200, r)
	return nil
}
