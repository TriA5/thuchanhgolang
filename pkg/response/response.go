package response

import (
	"net/http"

	pkgErrors "thuchanhgolang/pkg/errors"

	"github.com/gin-gonic/gin"
)

const (
	// DefaultErrorMessage is the default error message.
	DefaultErrorMessage = "Something went wrong"
	// ValidationErrorCode is the validation error code.
	ValidationErrorCode = 401
	// ValidationErrorMsg is the validation error message.
	ValidationErrorMsg = "Validation error"
)

// Resp is the response format.
type Resp struct {
	ErrorCode int    `json:"error_code"`
	Message   string `json:"message"`
	Data      any    `json:"data,omitempty"`
}

// NewOKResp returns a new OK response with the given data.
func NewOKResp(data any) Resp {
	return Resp{
		ErrorCode: 0,
		Message:   "Success",
		Data:      data,
	}
}

// NewUnauthorizedResp returns a new Unauthorized response with the given data.
func NewUnauthorizedResp() Resp {
	return Resp{
		ErrorCode: 401,
		Message:   "Unauthorized",
	}
}

// NewForbiddenResp returns a new Forbidden response
func NewForbiddenResp() Resp {
	return Resp{
		ErrorCode: 403,
		Message:   "Forbidden: You don't have permission to access this resource",
	}
}

// Ok returns a new OK response with the given data.
func OK(c *gin.Context, data any) {
	c.JSON(http.StatusOK, NewOKResp(data))
}

// Unauthorized returns a new Unauthorized response with the given data.
func Unauthorized(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, NewUnauthorizedResp())
}

// Forbidden returns a new Forbidden response
func Forbidden(c *gin.Context) {
	c.JSON(http.StatusForbidden, NewForbiddenResp())
}

func parseError(err error) (int, Resp) {
	switch parsedErr := err.(type) {
	case *pkgErrors.ValidationErrorCollector:
		return http.StatusBadRequest, Resp{
			ErrorCode: ValidationErrorCode,
			Message:   ValidationErrorMsg,
			Data:      parsedErr.Errors(),
		}
	case *pkgErrors.HTTPError:
		return http.StatusBadRequest, Resp{
			ErrorCode: parsedErr.Code,
			Message:   parsedErr.Message,
		}
	default:
		return http.StatusInternalServerError, Resp{
			ErrorCode: 500,
			Message:   DefaultErrorMessage,
		}
	}
}

// Error returns a new Error response with the given error.
func Error(c *gin.Context, err error) {
	c.JSON(parseError(err))
}

// ErrorMapping is a map of error to HTTPError.
type ErrorMapping map[error]*pkgErrors.HTTPError

// ErrorWithMap returns a new Error response with the given error.
func ErrorWithMap(c *gin.Context, err error, eMap ErrorMapping) {
	if httpErr, ok := eMap[err]; ok {
		Error(c, httpErr)
		return
	}

	Error(c, err)
}
