package user

import (
	"back/internal/adapters/middlewares"
	"back/internal/domain/user"
	"back/pkg/logger"
	"back/pkg/mysqlClient"
	"github.com/gin-gonic/gin"
)

const (
	getByIdPath         = "/users/:id"
	getFullInfoByIdPath = "/users/:id/profile"
	createPath          = "/users"
	activatePath        = "/users/:id/activate"
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
	e.POST(createPath, h.createUser())
	e.POST(activatePath, h.activateUser())

	authorized := e.Group("/")
	authorized.Use(middlewares.AuthMiddleware)
	authorized.GET(getByIdPath, h.getUserById())
	authorized.GET(getFullInfoByIdPath, h.getUserFullInfoById())
}
