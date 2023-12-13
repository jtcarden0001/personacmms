package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	eq "github.com/jtcarden0001/personacmms/projects/webapi/pkg/equipment"
	tp "github.com/jtcarden0001/personacmms/projects/webapi/pkg/types"
)

func Create(c *gin.Context) {
	var e tp.Equipment
	if err := c.BindJSON(&e); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	_, err := eq.Create(e.Title, e.Description)
	if err != nil {
		c.IndentedJSON(http.StatusCreated, e)
	} else {
		// TODO: revisit this error to make sure it is correct
		c.JSON(400, gin.H{"error": err.Error()})
	}
}

func GetAll(c *gin.Context) {
	equipment, err := eq.GetAll()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	} else {
		c.IndentedJSON(http.StatusOK, equipment)
	}
}
