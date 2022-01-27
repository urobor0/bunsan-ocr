package health

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// CheckHandler returns an HTTP handler to perform health checks.
func CheckHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "everything is ok!")
	}
}
