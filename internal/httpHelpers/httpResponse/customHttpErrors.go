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
		switch err.Error() {
		case "crypto/bcrypt: hashedPassword is not the hash of the given password":
			message = "Incorrect email or password."
			code = http.StatusBadRequest
		default:
			message = "Something went wrong. Please try later."
			code = http.StatusBadRequest
		}

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

func RequestErrCustomMessage(c *gin.Context, err error, message string) {
	c.JSON(400, gin.H{
		"message": message,
		"reason":  err.Error(),
	})
}
