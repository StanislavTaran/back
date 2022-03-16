package user

import (
	"back/internal/domain/user"
	"back/pkg/logger"
	"back/pkg/mysqlClient"
	"github.com/gin-gonic/gin"
)

const (
	signInPath = "/auth/signIn"

	getByIdPath  = "/users/:id"
	createPath   = "/users"
	activatePath = "/users/:id/activate"
)

type Handler struct {
	userService userService
	logger      logger.ILogger
}

func NewUserHandler(storage *mysqlClient.MySQLClient, logger logger.ILogger) *Handler {
	userStorage := user.NewUserStorage(storage)
	userService := user.NewUserService(userStorage)
	return &Handler{
		userService: userService,
		logger:      logger,
	}
}

func (h *Handler) Register(e *gin.Engine) {
	e.POST(signInPath, h.signIn())

	e.POST(createPath, h.createUser())
	e.GET(getByIdPath, h.getUserById())
	e.POST(activatePath, h.activateUser())
}
