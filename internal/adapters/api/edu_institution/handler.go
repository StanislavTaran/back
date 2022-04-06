package edu_institution

import (
	"back/internal/adapters/middlewares"
	mysqlEduInstitution "back/internal/adapters/mysql/education_institution"
	"back/internal/domain/education_institution"
	"back/pkg/logger"
	"back/pkg/mysqlClient"
	"github.com/gin-gonic/gin"
)

const (
	getListByNamePath = "/list"
)

type Handler struct {
	eduInstitutionService eduInstitutionService
	logger                logger.ILogger
}

func NewEduInstitutionHandler(client *mysqlClient.MySQLClient, logger logger.ILogger) *Handler {
	eduInstitutionStorage := mysqlEduInstitution.NewEducationInstitutionStorage(client)
	eduInstitutionService := education_institution.NewEducationInstitutionService(eduInstitutionStorage)
	return &Handler{
		eduInstitutionService: eduInstitutionService,
		logger:                logger,
	}
}

func (h *Handler) Register(e *gin.Engine) {
	authorized := e.Group("/edu-inst")
	authorized.Use(middlewares.AuthMiddleware)
	authorized.GET(getListByNamePath, h.getEduInstitutionsListByName())
}
