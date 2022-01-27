package ocr_job

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

var (
	_, b , _, _ = runtime.Caller(0)
	basePath = filepath.Join(filepath.Dir(b), "../../../../../")
	resourcesPath = fmt.Sprintf("%s/resources", basePath)
)

func TestCreateOCRJobHandler(t *testing.T) {
	fileName := "input-example.txt"
	gin.SetMode(gin.TestMode)
	r := gin.New()
	
	r.POST("/ocr-job", CreateOCRJobHandler())
	
	t.Run("given a valid request return 200", func(t *testing.T) {
		bodyBuf := &bytes.Buffer{}
		bodyWriter := multipart.NewWriter(bodyBuf)
		fileWriter, err := bodyWriter.CreateFormFile("file", "input")
		require.NoError(t, err)

		fh, err := os.Open(fmt.Sprintf("%s/%s", resourcesPath, fileName))
		require.NoError(t, err)

		_, err = io.Copy(fileWriter, fh)
		require.NoError(t, err)

		contentType := bodyWriter.FormDataContentType()
		bodyWriter.Close()

		req, err := http.NewRequest(http.MethodPost, "/ocr-job", bodyBuf)
		require.NoError(t, err)

		req.Header.Set("Content-Type", contentType)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusAccepted, res.StatusCode)
	})
	
	t.Run("given a invalid request return 400", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodPost, "/ocr-job", nil)
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	})
}
