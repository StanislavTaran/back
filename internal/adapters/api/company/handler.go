package company

import (
	"back/internal/adapters/middlewares"
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

func NewCompanyHandler(storage *mysqlClient.MySQLClient, logger logger.ILogger) *Handler {
	companyStorage := company.NewCompanyStorage(storage)
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
