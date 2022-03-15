package user

import (
	"back/internal/domain/user"
	"back/pkg/mysqlClient"
	"github.com/gin-gonic/gin"
)

const (
	getByIdPath = "/users/:id"
	createPath  = "/users"
)

type Handler struct {
	userService userService
}

func NewUserHandler(storage *mysqlClient.MySQLClient) *Handler {
	userStorage := user.NewUserStorage(storage)
	userService := user.NewUserService(userStorage)
	return &Handler{
		userService: userService,
	}
}

func (h *Handler) Register(e *gin.Engine) {
	e.POST(createPath, h.CreateUser())
	e.GET(getByIdPath, h.GetUserById())
}
