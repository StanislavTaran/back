package auth

import (
	"back/internal/adapters/middlewares"
	"back/internal/adapters/mysql/user"
	"back/internal/domain/auth"
	freecachepackage "back/pkg/cache/freecache"
	jwtpackage "back/pkg/jwt"
	"back/pkg/logger"
	"back/pkg/mysqlClient"
	"github.com/gin-gonic/gin"
)

const (
	signInPath       = "/auth/signIn"
	refreshTokenPath = "/auth/refresh"
	logOutPath       = "/auth/logout"
)

type Handler struct {
	authService authService
	logger      logger.ILogger
}

func NewAuthHandler(client *mysqlClient.MySQLClient, logger logger.ILogger) *Handler {
	userStorage := user.NewUserStorage(client)

	RTCache := freecachepackage.NewCacheRepo(100 * 1024 * 1024)
	helper := jwtpackage.NewHelper(RTCache, logger)
	authService := auth.NewAuthService(userStorage, helper)

	return &Handler{
		authService: authService,
		logger:      logger,
	}
}

func (h *Handler) Register(e *gin.Engine) {
	e.POST(signInPath, h.signIn())

	authorized := e.Group("/")
	authorized.Use(middlewares.AuthMiddleware)
	authorized.POST(refreshTokenPath, h.refreshToken())
	authorized.POST(logOutPath, h.logOut())
}
