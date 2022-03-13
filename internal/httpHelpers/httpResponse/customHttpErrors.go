package httpResponse

import "github.com/gin-gonic/gin"

func InternalErr(c *gin.Context, err error) {
	c.JSON(500, gin.H{
		"message": "Internal server error.",
		"reason":  err.Error(),
	})
}

func RequestErr(c *gin.Context, err error) {
	c.JSON(400, gin.H{
		"message": "Request error.",
		"reason":  err.Error(),
	})
}
