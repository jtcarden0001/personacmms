package gin

import "github.com/gin-gonic/gin"

func getStatus(err error, successCode int) int {
	if err != nil {
		return 500
	}
	return successCode
}

func getResponse(err error, data interface{}) interface{} {
	if err != nil {
		return gin.H{"error": err.Error()}
	}
	return data
}
