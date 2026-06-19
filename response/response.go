package response

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/Destiny-Peru/common/http/middleware"
)

type Payload struct {
	Success    bool      `json:"success"`
	StatusCode int       `json:"statusCode"`
	Message    string    `json:"message"`
	Data       any       `json:"data,omitempty"`
	Meta       any       `json:"meta,omitempty"`
	Error      any       `json:"error,omitempty"`
	RequestID  string    `json:"requestId,omitempty"`
	Path       string    `json:"path"`
	Timestamp  time.Time `json:"timestamp"`
}

type ErrorDetail struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message"`
}

func OK(c *gin.Context, data any) {
	Write(c, http.StatusOK, true, "Solicitud procesada correctamente.", data, nil, nil)
}

func OKWithMessage(c *gin.Context, message string, data any) {
	Write(c, http.StatusOK, true, message, data, nil, nil)
}

func Error(c *gin.Context, status int, message string) {
	Write(c, status, false, message, nil, nil, ErrorDetail{
		Message: message,
	})
}

func ErrorWithCode(c *gin.Context, status int, message string, code string) {
	Write(c, status, false, message, nil, nil, ErrorDetail{
		Code:    code,
		Message: message,
	})
}

func Write(c *gin.Context, status int, success bool, message string, data any, meta any, err any) {
	c.JSON(status, Payload{
		Success:    success,
		StatusCode: status,
		Message:    message,
		Data:       data,
		Meta:       meta,
		Error:      err,
		RequestID:  c.GetString(middleware.RequestIDKey),
		Path:       c.Request.URL.Path,
		Timestamp:  time.Now().UTC(),
	})
}
