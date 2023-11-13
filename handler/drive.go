package handler

import (
	"gdrive_uploader/pkg/drive_service"
	"gdrive_uploader/pkg/response"
	"github.com/gin-gonic/gin"
	"log"
)

func File(c *gin.Context) {
	srv, err := drive_service.NewService()
	if err != nil {
		log.Fatalf("Unable to retrieve Drive client: %v", err)
	}

	data, err := drive_service.ListFile(srv)
	if err != nil {
		c.JSON(200, response.NewResponse(false, err.Error(), nil))
	}
	c.JSON(200, response.NewResponse(true, "success", data))
}

func Upload(c *gin.Context) {
	srv, err := drive_service.NewService()
	if err != nil {
		log.Fatalf("Unable to retrieve Drive client: %v", err)
	}

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(200, response.NewResponse(false, "No file is received", nil))
	}

	blobFile, err := file.Open()
	if err != nil {
		c.JSON(200, response.NewResponse(false, err.Error(), nil))
	}

	err = drive_service.UploadFile(srv, blobFile, file.Filename)
	if err != nil {
		c.JSON(200, response.NewResponse(false, err.Error(), nil))
	}

	c.JSON(200, response.NewResponse(true, "success", nil))

}
