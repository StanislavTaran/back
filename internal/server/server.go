package server

import (
	"back/internal/adapters/api/user"
	"back/pkg/logger"
	"back/pkg/mysqlClient"
	"github.com/gin-gonic/gin"
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
	s.Engine.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	userHandler := user.NewUserHandler(s.storage, s.Logger)
	userHandler.Register(s.Engine)
}
