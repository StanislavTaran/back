package httpResponse

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ErrorByType(c *gin.Context, err error) {
	var message string
	var code int

	switch err {
	case sql.ErrNoRows:
		message = "No results found."
		code = http.StatusBadRequest
	case sql.ErrTxDone:
		message = "Operation failed. Please try later."
		code = http.StatusInternalServerError
	default:
		message = "Something went wrong. Please try later."
		code = http.StatusInternalServerError
	}

	c.JSON(code, gin.H{
		"message": message,
		"reason":  err.Error(),
	})
}

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
