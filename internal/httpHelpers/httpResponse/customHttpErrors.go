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

type ResponseError struct {
	Message string `json:"message"`
	Reason  string `json:"reason"`
}

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

	c.JSON(code, &ResponseError{
		Message: message,
		Reason:  err.Error(),
	})
}

func InternalErr(c *gin.Context, err error) {
	c.JSON(500, &ResponseError{
		Message: APP_ERR,
		Reason:  err.Error(),
	})
}

func RequestErr(c *gin.Context, err error) {
	c.JSON(400, &ResponseError{
		Message: REQ_ERR,
		Reason:  err.Error(),
	})
}

func RequestErrCustomMessage(c *gin.Context, err error, message string) {
	c.JSON(400, &ResponseError{
		Message: message,
		Reason:  err.Error(),
	})
}
