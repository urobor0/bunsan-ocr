package ocr_job

import (
	"bunsan-ocr/internal/ocr"
	"bunsan-ocr/internal/ocr/creating"
	"bunsan-ocr/kit/bus/command"
	"bunsan-ocr/kit/identifier"
	"bunsan-ocr/kit/projectpath"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

type CreateOCRJobRequest struct {
	File *multipart.FileHeader `form:"file" binding:"required"`
}

func CreateOCRJobHandler(bus command.Bus) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request CreateOCRJobRequest

		jobId, err := identifier.New()
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		if err := ctx.ShouldBind(&request); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "No file is received",
			})
			return
		}

		file, err := request.File.Open()
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		defer file.Close()

		fileContentType, err := getContentType(file)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		fileInputPath, err := saveFile(file, fmt.Sprintf("%s-input", jobId), filepath.Ext(request.File.Filename))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		err = bus.Dispatch(ctx, creating.NewJobCommand(
			jobId,
			fileInputPath,
			fileContentType,
		))

		if err != nil {
			switch {
			case errors.Is(err, ocr.ErrInvalidJobID):
				ctx.JSON(http.StatusBadRequest, err.Error())
				return
			default:
				ctx.JSON(http.StatusInternalServerError, err.Error())
				return
			}
		}

		ctx.JSON(http.StatusAccepted, gin.H{
			"input_path":   fileInputPath,
			"content_type": fileContentType,
			"message":      fmt.Sprintf("check the progress of your work at /ocr-job/%s", jobId),
		})
	}
}

func getContentType(file multipart.File) (string, error) {
	fileHeader := make([]byte, 512)

	// Copy the headers into the FileHeader buffer
	if _, err := file.Read(fileHeader); err != nil {
		return "", err
	}

	// set position back to start.
	if _, err := file.Seek(0, 0); err != nil {
		return "", err
	}

	return http.DetectContentType(fileHeader), nil
}

func saveFile(file multipart.File, fileName, extension string) (string, error) {
	attachmentFolder := fmt.Sprintf("%s/attachments", projectpath.RootDir())
	filePathDest := fmt.Sprintf("%s/%s%s", attachmentFolder, fileName, extension)

	// Create the uploads' folder if it doesn't
	// already exist
	if err := os.MkdirAll(attachmentFolder, os.ModePerm); err != nil {
		return "", err
	}

	// Create a new file in the uploads directory
	dst, err := os.Create(filePathDest)
	if err != nil {
		return "", err
	}

	defer dst.Close()

	// Copy the uploaded file to the filesystem
	// at the specified destination
	_, err = io.Copy(dst, file)
	if err != nil {
		return "", err
	}

	return filePathDest, err
}
