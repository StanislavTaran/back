package server

import (
	"back/internal/adapters/api/auth"
	"back/internal/adapters/api/company"
	"back/internal/adapters/api/edu_institution"
	"back/internal/adapters/api/user"
	"back/internal/adapters/api/user_company"
	"back/internal/adapters/api/user_education"
	"back/internal/adapters/middlewares"
	"back/pkg/logger"
	"back/pkg/mysqlClient"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "back/docs"
)

type Server struct {
	config  *Config
	Engine  *gin.Engine
	storage *mysqlClient.MySQLClient
	Logger  logger.ILogger
}

func NewServer(config *Config, logger logger.ILogger) *Server {
	return &Server{
		config: config,
		Logger: logger,
		Engine: gin.Default(),
	}
}

func (s *Server) Run() (err error) {
	err = s.configureMySQLStorage()
	if err != nil {
		return err
	}

	s.initRoutes()
	s.Logger.Debug("Routes mounted successfully.")

	err = s.Engine.Run(s.config.BindAddr)
	if err != nil {
		return err
	}

	s.Logger.Debug("Server started successfully")
	return nil
}

func (s *Server) configureMySQLStorage() error {
	storage := mysqlClient.NewMySQLClient(s.config.Mysql)

	err := storage.Open()
	if err != nil {
		return err
	}

	s.storage = storage
	return nil
}

func (s *Server) initRoutes() {
	s.Engine.Use(middlewares.CORSMiddleware())
	s.Engine.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	s.Engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	authHandler := auth.NewAuthHandler(s.storage, s.Logger)
	authHandler.Register(s.Engine)

	userHandler := user.NewUserHandler(s.storage, s.Logger)
	userHandler.Register(s.Engine)

	userCompanyHandler := user_company.NewUserCompanyHandler(s.storage, s.Logger)
	userCompanyHandler.Register(s.Engine)

	companyHandler := company.NewCompanyHandler(s.storage, s.Logger)
	companyHandler.Register(s.Engine)

	userEducationHandler := user_education.NewUserEducationHandler(s.storage, s.Logger)
	userEducationHandler.Register(s.Engine)

	eduInstitutionHandler := edu_institution.NewEduInstitutionHandler(s.storage, s.Logger)
	eduInstitutionHandler.Register(s.Engine)
}
