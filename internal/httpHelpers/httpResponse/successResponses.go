package httpResponse

import "github.com/gin-gonic/gin"

func SuccessOK(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "OK",
	})
}

func SuccessData(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{
		"data": data,
	})
}
