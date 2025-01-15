package response

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"
)

var errorStatusMap = map[error]int{
	domain.ErrNotFound:           http.StatusNotFound,
	domain.ErrConflict:           http.StatusConflict,
	domain.ErrInvalidParam:       http.StatusBadRequest,
	domain.ErrInvalidQueryParams: http.StatusBadRequest,
	domain.ErrInvalidToken:       http.StatusUnauthorized,
	domain.ErrExpiredToken:       http.StatusUnauthorized,
}

// ValidationError sends an error response for specific request validation errors
func ValidationError(ctx *gin.Context, err error) {
	HandleErrorWithStatus(ctx, http.StatusBadRequest, err)
}

// HandleError determines the status code of an error and returns a JSON response
func HandleError(ctx *gin.Context, err error) {
	statusCode, ok := errorStatusMap[err]
	if !ok {
		statusCode = http.StatusInternalServerError
		slog.Error("unhandled error", "error", err)
	}
	HandleErrorWithStatus(ctx, statusCode, err)
}

func HandleErrorWithStatus(ctx *gin.Context, statusCode int, err error) {
	slog.Error("HandleErrorWithStatus", "error", err)
	errMsgs := parseError(err)
	errRsp := newErrorResponse(errMsgs)
	ctx.JSON(statusCode, errRsp)
}

func parseError(err error) []string {
	var errMsgs []string

	var validationErrs validator.ValidationErrors
	var errorsAs = errors.As(err, &validationErrs)
	if errorsAs {
		for _, vErr := range validationErrs {
			errMsgs = append(errMsgs, vErr.Error())
		}
	} else {
		errMsgs = append(errMsgs, err.Error())
	}

	return errMsgs
}

// ErrorResponse represents the standard error response format
type ErrorResponse struct {
	Errors []string `json:"errors" example:"['Validation failed: field X is required', 'Invalid format for field Y']"`
}

func newErrorResponse(errMsgs []string) ErrorResponse {
	return ErrorResponse{
		Errors: errMsgs,
	}
}
