package api

import (
	"errors"
	"net/http"

	domain "github.com/champion19/Flighthours_backend/internal/domain/employee"
	"github.com/gin-gonic/gin"
)

var (
	ErrUnmarshalBody = errors.New("error unmarshal request body")
	ErrValidationUser = errors.New("error validation user: %w")
)

type WebError struct {
	Status int `json:"status"`
	Message string `json:"message"`
}

func (h handler) HandleError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, ErrUnmarshalBody):
		c.JSON(http.StatusBadRequest, WebError{
			Status: http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	case errors.Is(err, ErrValidationUser):
		c.JSON(http.StatusBadRequest, WebError{
			Status: http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	case errors.Is(err, domain.ErrUserCannotSave):
		c.JSON(http.StatusFailedDependency, WebError{
			Status: http.StatusFailedDependency,
			Message: err.Error(),
		})
		return
	case errors.Is(err, domain.ErrDuplicateUser):
		c.JSON(http.StatusAlreadyReported, WebError{
			Status: http.StatusAlreadyReported,
			Message: err.Error(),
		})
		return
	default:
		c.JSON(http.StatusInternalServerError, WebError{
			Status: http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}
}
