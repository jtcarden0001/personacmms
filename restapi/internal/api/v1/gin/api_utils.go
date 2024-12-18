package gin

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
)

func getStatus(err error, successCode int) int {
	if err == nil {
		return successCode
	}

	var appErr ae.AppError
	if errors.As(err, &appErr) {
		switch appErr.Code {
		case ae.CodeNotFound:
			return http.StatusNotFound
		}
	}

	return http.StatusInternalServerError
}

func getResponse(err error, data interface{}) interface{} {
	if err != nil {
		return gin.H{"error": err.Error()}
	}
	return data
}
