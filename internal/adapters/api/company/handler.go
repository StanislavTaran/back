package company

import (
	"back/internal/adapters/middlewares"
	companyMysql "back/internal/adapters/mysql/company"
	"back/internal/domain/company"
	"back/pkg/logger"
	"back/pkg/mysqlClient"
	"github.com/gin-gonic/gin"
)

const (
	getListByNamePath = "/list"
)

type Handler struct {
	companyService companyService
	logger         logger.ILogger
}

func NewCompanyHandler(client *mysqlClient.MySQLClient, logger logger.ILogger) *Handler {
	companyStorage := companyMysql.NewCompanyStorage(client)
	companyService := company.NewCompanyService(companyStorage)
	return &Handler{
		companyService: companyService,
		logger:         logger,
	}
}

func (h *Handler) Register(e *gin.Engine) {
	authorized := e.Group("/company")
	authorized.Use(middlewares.AuthMiddleware)
	authorized.GET(getListByNamePath, h.getCompaniesListByName())
}
