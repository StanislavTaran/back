package httpResponse

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	APP_ERR = "Something went wrong. Please try later."

	REQ_ERR                = "Request error"
	REQ_ERR_USER_NOT_FOUND = "User not found"

	VALIDATION_ERR = "Validation error"
)

func ErrorByType(c *gin.Context, err error) {
	var message string
	var code int

	switch err {
	case sql.ErrNoRows:
		message = "No results found."
		code = http.StatusBadRequest
	case sql.ErrTxDone:
		message = APP_ERR
		code = http.StatusInternalServerError
	default:
		switch err.Error() {
		case "crypto/bcrypt: hashedPassword is not the hash of the given password":
			message = "Incorrect email or password."
			code = http.StatusBadRequest
		default:
			message = APP_ERR
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
		"message": APP_ERR,
		"reason":  err.Error(),
	})
}

func RequestErr(c *gin.Context, err error) {
	c.JSON(400, gin.H{
		"message": REQ_ERR,
		"reason":  err.Error(),
	})
}

func RequestErrCustomMessage(c *gin.Context, err error, message string) {
	c.JSON(400, gin.H{
		"message": message,
		"reason":  err.Error(),
	})
}
