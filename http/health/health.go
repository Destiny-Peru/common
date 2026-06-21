package health

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Destiny-Peru/common/response"
)

type Checker interface {
	Name() string
	Check(ctx context.Context) error
}

type StatusData struct {
	Status string `json:"status"`
}

type Options struct {
	SuccessMessage string
	FailureMessage string
	ErrorCode      string
	Checkers       []Checker
}

func Handler(options Options) gin.HandlerFunc {
	successMessage := options.SuccessMessage
	if successMessage == "" {
		successMessage = "Service healthy."
	}

	failureMessage := options.FailureMessage
	if failureMessage == "" {
		failureMessage = "Service unavailable."
	}

	errorCode := options.ErrorCode
	if errorCode == "" {
		errorCode = "dependency_unavailable"
	}

	return func(c *gin.Context) {
		for _, checker := range options.Checkers {
			if err := checker.Check(c.Request.Context()); err != nil {
				errorMessage := err.Error()
				if checker.Name() != "" {
					errorMessage = checker.Name() + ": " + errorMessage
				}

				response.Write(
					c,
					http.StatusServiceUnavailable,
					false,
					failureMessage,
					nil,
					nil,
					response.ErrorDetail{
						Code:    errorCode,
						Message: errorMessage,
					},
				)

				return
			}
		}

		response.OKWithMessage(c, successMessage, StatusData{Status: "ok"})
	}
}

func Register(router gin.IRoutes, path string, options Options) {
	router.GET(path, Handler(options))
}
