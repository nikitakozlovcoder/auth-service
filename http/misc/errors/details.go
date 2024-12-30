package misc

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorDetails(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"error": err.Error(),
	})
}

func BadRequestDetails(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{
		"error": err.Error(),
	})
}
