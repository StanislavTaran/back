package user

import (
	"back/internal/domain/user"
	"back/pkg/mysqlClient"
	"github.com/gin-gonic/gin"
)

const (
	getByIdPath = "/user/:id"
)

type Handler struct {
	userService userService
}

func NewUserHandler(storage *mysqlClient.MySQLClient) *Handler {
	us := user.NewUserService(storage)
	return &Handler{
		userService: us,
	}
}

func (h *Handler) Register(e *gin.Engine) {
	e.GET(getByIdPath, h.GetById())
}
