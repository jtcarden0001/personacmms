package gin

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
)

var errorKey = "error"

func getStatus(err error, successCode int) int {
	if err == nil {
		return successCode
	}

	var appErr ae.AppError
	if errors.As(err, &appErr) {
		switch appErr.Code {

		case ae.CodeNotFound:
			return http.StatusNotFound

		case ae.CodeInvalid:
			return http.StatusBadRequest
		}

	}

	return http.StatusInternalServerError
}

func getResponse(err error, data interface{}) interface{} {
	if err != nil {
		return gin.H{errorKey: err.Error()}
	}
	return data
}
