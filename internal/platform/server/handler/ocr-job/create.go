package ocr_job

import (
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"net/http"
)

type CreateOCRJobRequest struct {
	File *multipart.FileHeader `form:"file" binding:"required"`
}

func CreateOCRJobHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request CreateOCRJobRequest

		if err := ctx.ShouldBind(&request); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "No file is received",
			})
			return
		}

		ctx.JSON(http.StatusAccepted, gin.H{})
	}
}