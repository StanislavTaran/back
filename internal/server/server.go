package server

import (
	"back/internal/adapters/api/user"
	"back/pkg/mysqlClient"
	"github.com/gin-gonic/gin"
)

type Server struct {
	config  *Config
	Engine  *gin.Engine
	storage *mysqlClient.MySQLClient
	//Logger logger.ILogger
}

func NewServer(config *Config) *Server {
	return &Server{
		config: config,
		//Logger: logger,
		Engine: gin.Default(),
	}
}

func (s *Server) Run() (err error) {
	err = s.configureMySQLStorage()
	if err != nil {
		return err
	}

	s.initRoutes()

	err = s.Engine.Run(s.config.BindAddr)
	if err != nil {
		return err
	}

	//s.Logger.Info("Server started successfully")
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

	userHandler := user.NewUserHandler(s.storage)
	userHandler.Register(s.Engine)
}
