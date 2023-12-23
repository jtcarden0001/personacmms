package gin

import "github.com/gin-gonic/gin"

func processAppError(c *gin.Context, err error) {
	//TODO: procees different errors (i.e. validation, not found, etc.)
	c.JSON(500, gin.H{"error": err.Error()})
}
