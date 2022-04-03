package user

import (
	"back/internal/adapters/middlewares"
	"back/internal/domain/user"
	"back/pkg/logger"
	"back/pkg/minioClient"
	"back/pkg/mysqlClient"
	"github.com/gin-gonic/gin"
)

const (
	getByIdPath         = "/users/:id"
	userAvatarPath      = "/users/:id/avatar"
	getFullInfoByIdPath = "/users/:id/profile"
	createPath          = "/users"
)

type Handler struct {
	userService userService
	logger      logger.ILogger
}

func NewUserHandler(client *mysqlClient.MySQLClient, fileStorageClient *minioClient.MinioClient, logger logger.ILogger) *Handler {
	userStorage := user.NewUserStorage(client)
	userFileStorage := user.NewUserFileStorage(fileStorageClient)
	userService := user.NewUserService(userStorage, userFileStorage)
	return &Handler{
		userService: userService,
		logger:      logger,
	}
}

func (h *Handler) Register(e *gin.Engine) {
	e.POST(createPath, h.createUser())

	authorized := e.Group("/")
	authorized.Use(middlewares.AuthMiddleware)
	authorized.GET(getByIdPath, h.getUserById())
	authorized.GET(getFullInfoByIdPath, h.getUserFullInfoById())
	authorized.POST(userAvatarPath, h.uploadAvatar())
}
