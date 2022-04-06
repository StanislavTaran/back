package user_company

import (
	"back/internal/adapters/middlewares"
	"back/internal/adapters/mysql/company"
	mysqlUserCompany "back/internal/adapters/mysql/user_company"
	"back/internal/domain/user_company"
	"back/pkg/logger"
	"back/pkg/mysqlClient"
	"github.com/gin-gonic/gin"
)

const (
	createPath = "/userCompany"
)

type Handler struct {
	userCompanyService userCompanyService
	logger             logger.ILogger
}

func NewUserCompanyHandler(storage *mysqlClient.MySQLClient, logger logger.ILogger) *Handler {
	companyStorage := company.NewCompanyStorage(storage)
	userCompanyStorage := mysqlUserCompany.NewUserCompanyStorage(storage)
	userCompanyService := user_company.NewUserCompanyService(userCompanyStorage, companyStorage)
	return &Handler{
		userCompanyService: userCompanyService,
		logger:             logger,
	}
}

func (h *Handler) Register(e *gin.Engine) {
	authorized := e.Group("/")
	authorized.Use(middlewares.AuthMiddleware)
	authorized.POST(createPath, h.createUserJobExperience())
}
